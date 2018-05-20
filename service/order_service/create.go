package order_service

import (
	"github.com/foxiswho/echo-go/models"
	"github.com/foxiswho/echo-go/module/sn"
	"github.com/foxiswho/echo-go/consts/order_consts"
	"time"
)

type CreateOrder struct {
	Order     *models.Order
	OrderExt  *models.OrderExt
	Consignee *models.OrderConsignee
	User      *models.User
	Goods     *models.Goods
}

func NewCreateOrder() *CreateOrder {
	return new(CreateOrder)
}

//默认配置
func (s *CreateOrder) defaultOption() {
	s.Order.TypeId = order_consts.PAY_ID_DEFAULT
	s.Order.GmtCreate = time.Now()
	s.Order.GmtModified = s.Order.GmtCreate
}

func (s *CreateOrder) formatOrder() {
	s.Order.OrderNo = sn.MakeOrderNo()

}
