package example

import (
	"net/http"

	. "github.com/foxiswho/echo-go/conf"
	"github.com/foxiswho/echo-go/middleware/opentracing"
	"github.com/foxiswho/echo-go/module/cache"
	"github.com/foxiswho/echo-go/module/session"
	"github.com/foxiswho/echo-go/router/base"
	"github.com/foxiswho/echo-go/router/example/api"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
)

// -----
// API RoutersApi
// -----
func RoutersApi() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Context自定义
	e.Use(base.NewBaseContext())

	// Customization
	if Conf.ReleaseMode {
		e.Debug = false
	}
	e.Logger.SetPrefix("api")
	e.Logger.SetLevel(GetLogLvl())

	// Session
	e.Use(session.Session())

	// OpenTracing
	if !Conf.Opentracing.Disable {
		e.Use(opentracing.OpenTracing("api"))
	}

	// CSRF
	e.Use(mw.CSRFWithConfig(mw.CSRFConfig{
		TokenLookup: "form:X-XSRF-TOKEN",
	}))

	// Gzip
	e.Use(mw.GzipWithConfig(mw.GzipConfig{
		Level: 5,
	}))

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Static("/favicon.ico", "./assets/img/favicon.ico")

	// Cache
	e.Use(cache.Cache())

	// e.Use(ec.SiteCache(ec.NewMemcachedStore([]string{conf.MEMCACHED_SERVER}, time.Hour), time.Minute))
	// e.GET("/user_service/:id", ec.CachePage(ec.NewMemcachedStore([]string{conf.MEMCACHED_SERVER}, time.Hour), time.Minute, UserHandler))

	// RoutersApi
	//e.GET("/login", UserLoginHandler)
	//e.GET("/register", UserRegisterHandler)

	// Unauthenticated route
	e.GET("/", accessible)
	json := e.Group("/json")
	{
		json.GET("/jsonp", base.Handler(api.JsonpHandler))
	}
	// JWT
	j := e.Group("/jwt")
	{
		// Login route
		j.POST("/login", base.Handler(api.JwtLoginPostHandler))
		i := j.Group("/restricted")
		{
			// Configure middleware with the custom claims type
			config := mw.JWTConfig{
				Claims:     &api.JwtCustomClaims{},
				SigningKey: []byte(Conf.SessionSecretKey),
			}
			i.Use(mw.JWTWithConfig(config))
			//j.Use(mw.JWTWithConfig(mw.JWTConfig{
			//	SigningKey:  []byte(Conf.SessionSecretKey),
			//	//ContextKey:  "_user",
			//	//TokenLookup: "header:" + echo.HeaderAuthorization,
			//}))
			i.GET("/xx", api.JwtApiHandler)
		}

		//curl http://echo.api.localhost:8080/restricted/user -H "Authorization: Bearer XXX"
		//r.GET("/user_service", UserHandler)
	}
	// JWT
	r := e.Group("/jwt2")
	{
		r.Use(mw.JWTWithConfig(mw.JWTConfig{
			SigningKey:  []byte("secret"),
			ContextKey:  "_user",
			TokenLookup: "header:" + echo.HeaderAuthorization,
		}))

		r.GET("/login-in", base.Handler(api.JwtTesterApiHandler))
	}

	//post := r.Group("/post")
	//{
	//	post.GET("/save", PostSaveHandler)
	//	post.GET("/id/:id", PostHandler)
	//	post.GET("/:userId/p/:p/s/:s", PostsHandler)
	//}

	return e
}
func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}
