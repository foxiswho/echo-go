package rbac

import (
	"github.com/foxiswho/shop-go/router/base"
)

func IndexHandler(c *base.BaseContext) error {
	//a := authadapter.NewAdapter("mysql", "")
	//e := casbin.NewEnforcer("template/casbin/rbac_model.conf")
	// Load the policy from DB.
	//e.LoadPolicy()
	//fmt.Println("e.Enforce", e.Enforce("1", "/admin/rdbc/index", "GET"))
	//e.AddPolicy("1","/admin/rdbc/*","GET")
	// Save the policy back to DB.
	//e.SavePolicy()

	c.Set("tmpl", "example/admin/rbac/index")
	c.Set("data", map[string]interface{}{
		"title": "rbac",
	})

	return nil
}
