package test

import (
	"net/http"
	"github.com/foxiswho/shop-go/router/base"
	"github.com/foxiswho/shop-go/module/auth"
)

func LogoutHandler(c *base.BaseContext) error {
	session := c.Session()
	a := c.Auth()
	auth.Logout(session, a.User)

	redirect := c.QueryParam(auth.RedirectParam)
	if redirect == "" {
		redirect = "/"
	}

	c.Redirect(http.StatusMovedPermanently, redirect)

	return nil
}
