package opentracing

import (
	"fmt"
	"net"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/opentracing/opentracing-go"
	"sourcegraph.com/sourcegraph/appdash"
	"github.com/uber/jaeger-lib/metrics"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	appdashot "sourcegraph.com/sourcegraph/appdash/opentracing"
	"sourcegraph.com/sourcegraph/appdash/traceapp"
	"io"
)

const (
	DefaultKey = "github.com/foxiswho/shop-go/middleware/opentracing"

	TracerTypeJaeger = "jaeger"

	TracerTypeAppdash = "appdash"
)

type Configuration struct {
	Disabled bool
	Type     string
}

func (c Configuration) InitGlobalTracer(options ...Option) io.Closer {
	if c.Disabled {
		return nil
	} else {
		opts := applyOptions(c.Type, options...)

		switch c.Type {
		case TracerTypeAppdash:
			initGlobalTracer_Appdash(opts.Address)
			return nil
		case TracerTypeJaeger:
			return initGlobalTracer_Jaeger(opts.ServiceName, opts.Address)
		default:
			return nil
		}
	}
}

func initGlobalTracer_Jaeger(serviceName, addr string) io.Closer {
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := &jaegerLogger{}
	jMetricsFactory := metrics.NullFactory

	metricsFactory := metrics.NewLocalFactory(0)
	metrics := jaeger.NewMetrics(metricsFactory, nil)

	sender, err := jaeger.NewUDPTransport(addr, 0)
	if err != nil {
		log.Printf("could not initialize jaeger sender: %s", err.Error())
		return nil
	}

	repoter := jaeger.NewRemoteReporter(sender, jaeger.ReporterOptions.Metrics(metrics))

	// Initialize tracer with a logger and a metrics factory
	closer, err := cfg.InitGlobalTracer(
		serviceName,
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
		jaegercfg.Reporter(repoter),
	)

	if err != nil {
		log.Printf("could not initialize jaeger tracer: %s", err.Error())
		return nil
	}
	//defer closer.Close()
	return closer
}

func initGlobalTracer_Appdash(addr string) {
	// OpenTracing
	// glide get github.com/opentracing/opentracing-go
	// glide get sourcegraph.com/sourcegraph/appdash
	// glide get github.com/gogo/protobuf
	store := appdash.NewMemoryStore()

	// Listen on any available TCP port locally.
	l, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 0})
	if err != nil {
		log.Panic(err)
	}
	collectorPort := l.Addr().(*net.TCPAddr).Port
	collectorAdd := fmt.Sprintf(":%d", collectorPort)

	// Start an Appdash collection server that will listen for spans and
	// annotations and add them to the local collector (stored in-memory).
	cs := appdash.NewServer(l, appdash.NewLocalCollector(store))
	go cs.Start()

	// Print the URL at which the web UI will be running.
	appdashURL, err := url.Parse(addr)
	if err != nil {
		errStr := fmt.Sprintf("error appdash url parsing %s: %s", addr, err)
		log.Panic(errStr)
	}
	log.Debugf("to see your traces, go to %s/traces\n", appdashURL)

	// Start the web UI in a separate goroutine.
	tapp, err := traceapp.New(nil, appdashURL)
	if err != nil {
		log.Panic(err)
	}
	tapp.Store = store
	tapp.Queryer = store
	go func() {
		log.Fatal(http.ListenAndServe(":"+appdashURL.Port(), tapp))
	}()

	tracer := appdashot.NewTracer(appdash.NewRemoteCollector(collectorAdd))
	opentracing.InitGlobalTracer(tracer)
}

func OpenTracing(comp string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var span opentracing.Span
			opName := comp + ":" + c.Request().URL.Path
			// 监测Header中是否有Trace信息
			wireContext, err := opentracing.GlobalTracer().Extract(
				opentracing.TextMap,
				opentracing.HTTPHeadersCarrier(c.Request().Header))
			if err != nil {
				// 启动新Span
				span = opentracing.StartSpan(opName)
			} else {
				log.Debugf("opentracing span child!")
				span = opentracing.StartSpan(opName, opentracing.ChildOf(wireContext))
			}

			defer span.Finish()
			c.Set(DefaultKey, span)

			span.SetTag("component", comp)
			span.SetTag("span.kind", "server")
			span.SetTag("http.url", c.Request().Host+c.Request().RequestURI)
			span.SetTag("http.method", c.Request().Method)

			if err := next(c); err != nil {
				span.SetTag("error", true)
				c.Error(err)
			}

			span.SetTag("error", false)
			span.SetTag("http.status_code", c.Response().Status)

			return nil
		}
	}
}

func Default(c echo.Context) opentracing.Span {
	ot := c.Get(DefaultKey)
	if ot == nil {
		return nil
	}
	return c.Get(DefaultKey).(opentracing.Span)
}

type jaegerLogger struct{}

func (l *jaegerLogger) Error(msg string) {
	log.Debugf("ERROR: %s", msg)
}

// Infof logs a message at info priority
func (l *jaegerLogger) Infof(msg string, args ...interface{}) {
	log.Debugf(msg, args...)
}
