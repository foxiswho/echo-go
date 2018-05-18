package admin

import (
	"github.com/foxiswho/shop-go/router/base"
	"github.com/casbin/casbin"
	"github.com/foxiswho/shop-go/middleware/authadapter"
	"fmt"
)

func adminAuth(c *base.BaseContext) error {
	// Initialize a Xorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	a := authadapter.NewAdapter("mysql", "")
	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	e := casbin.NewEnforcer("template/casbin/rbac_model.conf", a)

	// Load the policy from DB.
	//e.LoadPolicy()
	//e.AddRoleForUser()
	// Check the permission.
	fmt.Println("e.Enforce", e.Enforce("alice", "data1", "read"))

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	//e.SavePolicy()
	return nil
}

func CheckAuthCasbin() {
	a := authadapter.NewAdapter("mysql", "")
	e := casbin.NewEnforcer("template/casbin/rbac_model.conf", a)
	fmt.Println("e.Enforce", e.Enforce("alice", "data1", "read"))
}
