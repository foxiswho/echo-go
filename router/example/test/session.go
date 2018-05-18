package test

import (
	"github.com/foxiswho/echo-go/router/base"
	"fmt"
)

//
type Session struct {

}

func NewSession() *Session{
	return  new(Session)
}

func (x *Session) IndexHandler(c *base.BaseContext) error {
	c.Session().Set("SSSSSSS","asldfjlksajdflkasjdflkjd")

	test:=c.Session().Get("SSSSSSS")
	fmt.Println("session=》SSSSSSS",test)

	c.Set("tmpl", "example/test/session")
	c.Set("data", map[string]interface{}{
		"title": "测试 COOIE",
		"test":  test,
	})

	return nil
}