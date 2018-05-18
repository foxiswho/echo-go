package admin

import (
	"github.com/casbin/casbin"
	"github.com/labstack/echo"
	casbinmw "github.com/labstack/echo-contrib/casbin"
)

func main() {
	ce := casbin.NewEnforcer("auth_model.conf", "")
	ce.AddRoleForUser("alice", "admin")
	e := echo.New()
	e.Use(casbinmw.Middleware(ce))
	e.Logger.Fatal(e.Start(":1323"))
}
