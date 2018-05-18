package test

import (
	"github.com/foxiswho/shop-go/router/base"
)

func JwtLoginHandler(c *base.BaseContext) error {
	c.Set("tmpl", "example/test/jwt_login")
	c.Set("data", map[string]interface{}{
		"title": "JWT 接口测试",
	})

	return nil
}
