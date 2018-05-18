package index

import (
	"github.com/foxiswho/echo-go/router/base"
)

func AboutHandler(c *base.BaseContext) error {
	c.Set("tmpl", "web/index/about")
	c.Set("data", map[string]interface{}{
		"title": "About",
	})

	return nil
}
