package rbac

import (
	"github.com/foxiswho/echo-go/router/base"
	"github.com/foxiswho/echo-go/middleware/authadapter"
	"github.com/casbin/casbin"
	"fmt"
)

func CreateHandler(c *base.BaseContext) error {
	a := authadapter.NewAdapter("mysql", "")
	e := casbin.NewEnforcer("template/casbin/rbac_model.conf", a)
	//添加
	e.AddPolicy("1", "/admin/rdbc/*", "GET")
	// Save the policy back to DB.
	e.SavePolicy()

	fmt.Println("e.Enforce", e.Enforce("1", "/admin/rdbc/index", "GET"))

	c.Set("tmpl", "example/admin/rbac/index")
	c.Set("data", map[string]interface{}{
		"title": "rbac",
	})

	return nil
}
