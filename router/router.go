package router

import (
	"context"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	. "github.com/foxiswho/shop-go/conf"
	"github.com/foxiswho/shop-go/middleware/opentracing"
	"github.com/foxiswho/shop-go/router/example"
	"github.com/foxiswho/shop-go/router/example/socket"
	"github.com/foxiswho/shop-go/router/web"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func InitRoutes() map[string]*Host {
	// Hosts
	hosts := make(map[string]*Host)

	hosts[Conf.Server.DomainWeb] = &Host{web.Routers()}
	hosts[Conf.Server.DomainApi] = &Host{example.RoutersApi()}
	hosts[Conf.Server.DomainSocket] = &Host{socket.Routers()}

	return hosts
}

// 子域名部署
func RunSubdomains(confFilePath string) {
	// 配置初始化
	if err := InitConfig(confFilePath); err != nil {
		log.Panic(err)
	}

	// 全局日志级别
	log.SetLevel(GetLogLvl())

	// Server
	e := echo.New()
	e.Pre(mw.RemoveTrailingSlash())

	// OpenTracing
	otCtf := opentracing.Configuration{
		Disabled: Conf.Opentracing.Disable,
		Type:     Conf.Opentracing.Type,
	}
	if closer := otCtf.InitGlobalTracer(
		opentracing.ServiceName(Conf.Opentracing.ServiceName),
		opentracing.Address(Conf.Opentracing.Address),
	); closer != nil {
		defer closer.Close()
	}

	// 日志级别
	e.Logger.SetLevel(GetLogLvl())

	// Secure, XSS/CSS HSTS
	e.Use(mw.SecureWithConfig(mw.DefaultSecureConfig))
	mw.MethodOverride()

	// CORS
	e.Use(mw.CORSWithConfig(mw.CORSConfig{
		AllowOrigins: []string{"http://" + Conf.Server.DomainWeb, "http://" + Conf.Server.DomainApi},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAuthorization},
	}))

	hosts := InitRoutes()
	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()

		u, _err := url.Parse(c.Scheme() + "://" + req.Host)
		if _err != nil {
			e.Logger.Errorf("Request URL parse error:%v", _err)
		}

		host := hosts[u.Hostname()]
		if host == nil {
			e.Logger.Info("Host not found")
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}

		return
	})

	if !Conf.Server.Graceful {
		e.Logger.Fatal(e.Start(Conf.Server.Addr))
	} else {
		// Graceful Shutdown
		// Start server
		go func() {
			if err := e.Start(Conf.Server.Addr); err != nil {
				e.Logger.Errorf("Shutting down the server with error:%v", err)
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 10 seconds.
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}
}
