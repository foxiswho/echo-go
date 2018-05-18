package api

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"net/http"
	"github.com/labstack/echo"
	"github.com/foxiswho/shop-go/router/base"
	"fmt"
	"github.com/foxiswho/shop-go/router/example/test"
	"github.com/foxiswho/shop-go/service/example_service"
	"github.com/foxiswho/shop-go/module/log"
	"github.com/foxiswho/shop-go/conf"
)

// jwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func JwtLoginPostHandler(c *base.BaseContext) error {
	c.Response().Header().Del("Access-Control-Allow-Origin")
	c.Response().Header().Add("Access-Control-Allow-Origin","*")
	var form test.LoginForm
	if err := c.Bind(&form); err == nil {
		fmt.Println("form",form)
		u := example_service.GetUserByNicknamePwd(form.Nickname, form.Password)
		fmt.Println("db=>u")
		fmt.Println("db=>u")
		fmt.Println("db=>u")
		fmt.Println("db=>u")
		fmt.Println("db=>u")
		fmt.Println("db=>u",u)
		if u != nil {
			// Set custom claims
			claims := &JwtCustomClaims{
				form.Nickname,
				true,
				jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
				},
			}

			// Create token with claims
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			// Generate encoded token and send it as response.
			t, err := token.SignedString([]byte(conf.Conf.SessionSecretKey))
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, echo.Map{
				"token": t,
			})
		} else {
			return c.JSON(http.StatusOK, echo.Map{
				"message": "用户不存在",
			})
		}
	} else {
		params, _ := c.FormParams()
		log.Debugf("Login form params: %v", params)
		log.Debugf("Login form bind Error: %v", err)
		return c.JSON(http.StatusOK, echo.Map{
			"message": "错误",
		})
	}
	return echo.ErrUnauthorized
}
