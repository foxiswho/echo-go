package web

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	"github.com/foxiswho/echo-go/middleware/captcha"
	"github.com/foxiswho/echo-go/middleware/staticbin"

	"github.com/foxiswho/echo-go/assets"
	. "github.com/foxiswho/echo-go/conf"
	"github.com/foxiswho/echo-go/middleware/opentracing"
	"github.com/foxiswho/echo-go/module/auth"
	"github.com/foxiswho/echo-go/module/cache"
	"github.com/foxiswho/echo-go/module/render"
	"github.com/foxiswho/echo-go/module/session"
	serviceAuth "github.com/foxiswho/echo-go/service/user_service/auth"
	serviceAdminAuth "github.com/foxiswho/echo-go/service/admin_service/auth"
	web_index "github.com/foxiswho/echo-go/router/web/index"
	web_test "github.com/foxiswho/echo-go/router/example/test"
	example_admin "github.com/foxiswho/echo-go/router/example/admin"
	"github.com/foxiswho/echo-go/router/base"
	"github.com/foxiswho/echo-go/router/example/api"
	"github.com/foxiswho/echo-go/middleware/authadapter"
	"github.com/casbin/casbin"
	auth_casbin "github.com/foxiswho/echo-go/middleware/auth"
	rbac2 "github.com/foxiswho/echo-go/router/example/admin/rbac"
	"github.com/foxiswho/echo-go/router/web/design"
)

//---------
// Website Routers
//---------
func Routers() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Context自定义
	e.Use(base.NewBaseContext())
	// Customization
	if Conf.ReleaseMode {
		e.Debug = false
	}
	e.Logger.SetPrefix("web")
	e.Logger.SetLevel(GetLogLvl())

	// Session
	e.Use(session.Session())

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// 验证码，优先于静态资源
	e.Use(captcha.Captcha(captcha.Config{
		CaptchaPath: "/captcha/",
		SkipLogging: true,
	}))

	// 静态资源
	switch Conf.Static.Type {
	case BINDATA:
		e.Use(staticbin.Static(assets.Asset, staticbin.Options{
			Dir:         "/",
			SkipLogging: true,
		}))
	default:
		e.Static("/assets", "./assets")
	}

	// Gzip，在验证码、静态资源之后
	// 验证码、静态资源使用http.ServeContent()，与Gzip有冲突，Nginx报错，验证码无法访问
	e.Use(mw.GzipWithConfig(mw.GzipConfig{
		Level: 5,
	}))

	// OpenTracing
	if !Conf.Opentracing.Disable {
		e.Use(opentracing.OpenTracing("web"))
	}
	////////////////////////////
	j := e.Group("/jwt")
	{
		j.POST("/login", base.Handler(api.JwtLoginPostHandler))
		i := j.Group("/restricted")
		{
			config := mw.JWTConfig{
				Claims:     &api.JwtCustomClaims{},
				SigningKey: []byte(Conf.SessionSecretKey),
			}
			i.Use(mw.JWTWithConfig(config))
			i.GET("/xx", api.JwtApiHandler)
		}
	}
	////////////////////////////
	// CSRF
	e.Use(mw.CSRFWithConfig(mw.CSRFConfig{
		ContextKey:  "_csrf",
		TokenLookup: "form:_csrf",
	}))

	// 模板
	e.Renderer = render.LoadTemplates()
	e.Use(render.Render())

	// Cache
	e.Use(cache.Cache())
	// Auth
	//e.Use(auth.New(model.GenerateAnonymousUser))
	//e.Use(auth.New(serviceAuth.GenerateAnonymousUser))
	// Routers
	index := e.Group("")
	{
		index.Use(auth.New(serviceAuth.GenerateAnonymousUser))
		index.GET("/", base.Handler(web_index.HomeHandler))
		//
		about := index.Group("/about")
		about.Use(auth.LoginRequired())
		{
			about.GET("", base.Handler(web_index.AboutHandler))
		}
		//
		test := index.Group("/example/test")
		{
			test.GET("/jwt/tester", base.Handler(web_test.JwtTesterHandler))
			test.GET("/jwt/login", base.Handler(web_test.JwtLoginHandler))
			//test.POST("/jwt/login", base.Handler(web_test.JwtLoginPostHandler))
			test.GET("/ws", base.Handler(web_test.WsHandler))
			test.GET("/cache", base.Handler(web_test.CacheHandler))
			test.GET("/cookie", base.Handler(web_test.NewCookie().IndexHandler))
			test.GET("/session", base.Handler(web_test.NewSession().IndexHandler))
			test.GET("/orm", base.Handler(web_test.NewOrm().IndexHandler))
			test.GET("/login", base.Handler(web_test.LoginHandler))
			test.POST("/login", base.Handler(web_test.LoginPostHandler))
			test.GET("/logout", base.Handler(web_test.LogoutHandler))
			test.GET("/register", base.Handler(web_test.RegisterHandler))
			test.POST("/register", base.Handler(web_test.RegisterPostHandler))
			user := test.Group("/user_service")
			user.Use(auth.LoginRequired())
			{
				user.GET("/:id", base.Handler(web_test.UserHandler))
			}
			test.GET("/upload", base.Handler(web_test.NewUpload().UploadIndex))
			test.POST("/upload", base.Handler(web_test.UploadPostIndex))
			test.POST("/upload-more", base.Handler(web_test.UploadMorePostIndex))
			test.POST("/upload-db", base.Handler(web_test.UploadDbHandler))
			test.GET("/jsonp", base.Handler(web_test.JsonpIndexHandler))
		}
	}
	////////////////////////////
	/////admin
	admin_login := e.Group("/admin_login")
	{
		admin_login.Use(auth.New(serviceAdminAuth.GenerateAnonymousUser))
		admin_login.GET("/", base.Handler(example_admin.DefaultHandler))
		admin_login.GET("/login", base.Handler(example_admin.LoginHandler))
		admin_login.POST("/login", base.Handler(example_admin.LoginPostHandler))
		admin_login.GET("/logout", base.Handler(example_admin.LogoutHandler))
	}
	admin := e.Group("/admin")
	{
		admin.Use(auth.New(serviceAdminAuth.GenerateAnonymousUser))
		admin.GET("", base.Handler(example_admin.IndexHandler))
		rbac := admin.Group("/rbac")
		{
			//数据库驱动
			a := authadapter.NewAdapter("mysql", "")
			//加载 过滤条件
			ce := casbin.NewEnforcer("template/casbin/rbac_model.conf",a)
			//从数据库加载到内存中
			ce.LoadPolicy()
			//中间件
			rbac.Use(auth_casbin.Middleware(ce))
			rbac.GET("/index", base.Handler(rbac2.IndexHandler))
		}
		//设计
		des := admin.Group("/design")
		{
			//根据数据库生成 service
			des.GET("/service", base.Handler(design.ServiceMakeHandler))
		}
	}
	return e
}
