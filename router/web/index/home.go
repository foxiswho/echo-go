package index

import (
	"net/http"
	"github.com/opentracing/opentracing-go"
	"fmt"
	//"github.com/foxiswho/echo-go/model"
	"github.com/foxiswho/echo-go/module/model"
	"github.com/foxiswho/echo-go/module/log"
	. "github.com/foxiswho/echo-go/conf"
	sauth "github.com/foxiswho/echo-go/service/user_service/auth"
	"github.com/foxiswho/echo-go/router/base"
)

func HomeHandler(c *base.BaseContext) error {
	// OpenTracing层级监控示例，API层通过中间件已支持
	span := c.OpenTracingSpan()
	if span != nil {
		// Since we have to inject our span into the HTTP headers, we create a request
		asyncReq, _ := http.NewRequest("GET", "http://"+Conf.Server.DomainApi+"/login", nil)
		// Inject the span context into the header
		err := span.Tracer().Inject(span.Context(),
			opentracing.TextMap,
			opentracing.HTTPHeadersCarrier(asyncReq.Header))
		if err != nil {
			log.Debugf("Could not inject span context into header: %v", err)
		}
		go func() {
			if _, err := http.DefaultClient.Do(asyncReq); err != nil {
				span.SetTag("error", true)
				span.LogEvent(fmt.Sprintf("GET /login error: %v", err))
			}
		}()
	} else {
		log.Debugf("opentracing span nil")
	}

	User := new(sauth.User)
	User.Model = model.Model{Context: c}
	User.Id = 1
	User.TraceGetUserById(1)

	c.Set("tmpl", "web/index/home")
	c.Set("data", map[string]interface{}{
		"title": "Home",
	})

	return nil
}
