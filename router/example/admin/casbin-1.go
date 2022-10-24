package admin

import (
	"github.com/casbin/casbin"
	casbinmw "github.com/labstack/echo-contrib/casbin"
	"github.com/labstack/echo/v4"
)

func main() {
	ce := casbin.NewEnforcer("auth_model.conf", "")
	ce.AddRoleForUser("alice", "admin")
	e := echo.New()
	e.Use(casbinmw.Middleware(ce))
	e.Logger.Fatal(e.Start(":1323"))
}
