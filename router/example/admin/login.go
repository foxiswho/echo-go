package admin

import (
	"net/http"
	"github.com/foxiswho/shop-go/router/base"
	"github.com/foxiswho/shop-go/module/auth"
	"github.com/foxiswho/shop-go/module/log"
	userService "github.com/foxiswho/shop-go/service/example_service"
	"fmt"
)

type LoginForm struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func LoginHandler(c *base.BaseContext) error {
	//redirect := c.QueryParam(auth.RedirectParam)
	//
	//a := c.Auth()
	//if a.Admin.IsAuthenticated() {
	//	if redirect == "" {
	//		redirect = "/"
	//	}
	//	c.Redirect(http.StatusMovedPermanently, redirect)
	//	return nil
	//}
	//CheckAuthCasbin()

	c.Set("tmpl", "example/admin/login")
	c.Set("data", map[string]interface{}{
		"title":         "Login",
		//"redirectParam": auth.RedirectParam,
		//"redirect":      redirect,
	})

	return nil
}

func LoginPostHandler(c *base.BaseContext) error {
	loginURL := c.Request().RequestURI

	redirect := c.QueryParam(auth.RedirectParam)
	if redirect == "" {
		redirect = "/admin"
	}

	a := c.Auth()
	if a.User.IsAuthenticated() {
		fmt.Println("已经验证过了")
		c.Redirect(http.StatusMovedPermanently, redirect)
		return nil
	}

	var form LoginForm
	if err := c.Bind(&form); err == nil {
		fmt.Println("form",form)
		u := userService.NewAdminService().GetUserByNicknamePwd(form.UserName, form.Password)
		fmt.Println("db=>u",u)
		if u != nil {
			session := c.Session()
			err := auth.AuthenticateSession(session, u)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			c.Redirect(http.StatusMovedPermanently, redirect)
			return nil
		} else {
			c.Redirect(http.StatusMovedPermanently, loginURL)
			return nil
		}
	} else {
		params, _ := c.FormParams()
		log.Debugf("Login form params: %v", params)
		log.Debugf("Login form bind Error: %v", err)
		c.Redirect(http.StatusMovedPermanently, loginURL)
		return nil
	}
}
