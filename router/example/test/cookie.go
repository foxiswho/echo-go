package test

import (
	"github.com/foxiswho/shop-go/router/base"
	"fmt"
	"time"
)

type cookie struct {

}

func NewCookie() *cookie {
	return new(cookie)
}

func (x *cookie) IndexHandler(c *base.BaseContext) error {
	c.CookieSet("test", time.Now().String()+"=="+time.Now().String())
	value, err := c.CookieGet("test")
	fmt.Println("err", err)
	fmt.Println("cookie value=>", value)
	c.CookieSet("test2", time.Now().String())

	value2, err := c.CookieGet("test2")
	fmt.Println("err", err)
	fmt.Println("cookie value2=>", value2)




	c.Set("tmpl", "example/test/cookie")
	c.Set("data", map[string]interface{}{
		"title": "测试 COOIE",
		"test":  value,
		"test2": value2,
	})

	return nil
}
