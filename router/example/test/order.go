package test

import (
	"github.com/foxiswho/echo-go/router/base"
	"net/http"
	"github.com/foxiswho/echo-go/service/order_service"
)

type Order struct {
}

func NewOrder() *Order {
	return new(Order)
}

func (x *Order) IndexHandler(c *base.BaseContext) error {
	//test
	order_service.CreateOrderFormatFromCartT()
	c.HTML(http.StatusOK, "order 创建成功")
	return nil
}
