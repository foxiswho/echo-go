package test

import (
	"github.com/foxiswho/echo-go/router/base"
)

func WsHandler(c *base.BaseContext) error {
	c.Set("tmpl", "example/test/ws")
	c.Set("data", map[string]interface{}{
		"title": "Web Socket",
	})
	return nil
}
