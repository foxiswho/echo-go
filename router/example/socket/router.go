package socket

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/foxiswho/echo-go/module/auth"
	"github.com/foxiswho/echo-go/module/cache"
	"github.com/foxiswho/echo-go/module/render"
	"github.com/foxiswho/echo-go/module/session"
	authService "github.com/foxiswho/echo-go/service/user_service/auth"
)

func Routers() *echo.Echo {
	e := echo.New()

	// Session
	e.Use(session.Session())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 模板
	e.Renderer = render.LoadTemplates()
	e.Use(render.Render())

	// Cache
	e.Use(cache.Cache())

	// Auth
	e.Use(auth.New(authService.GenerateAnonymousUser))

	e.GET("/ws", socketHandler)

	return e
}
