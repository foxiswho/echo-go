package admin

import (
	"github.com/foxiswho/echo-go/router/base"
	"net/http"
)

func DefaultHandler(c *base.BaseContext) error {
	c.Redirect(http.StatusMovedPermanently, "/login")
	return nil
}
