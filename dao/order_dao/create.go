package order_dao

import (
	"github.com/foxiswho/echo-go/models"
	"github.com/foxiswho/echo-go/module/db"
	"github.com/foxiswho/echo-go/util"
	"github.com/foxiswho/echo-go/module/log"
)

type CreateOrder struct {
	*models.OrderCollectDate
}

func NewCreateOrder() *CreateOrder {
	return new(CreateOrder)
}

func (s *CreateOrder) SetOrderCollectDate(data *models.OrderCollectDate) {
	s.OrderCollectDate = data
}

func (s *CreateOrder) SetOrder(order *models.Order) {
	s.Order = order
}
func (s *CreateOrder) SetOrderExt(order_ext *models.OrderExt) {
	s.OrderExt = order_ext
}

func (s *CreateOrder) SetOrderConsignee(data *models.OrderConsignee) {
	s.OrderConsignee = data
}

func (s *CreateOrder) SetOrderGoods(data []*models.OrderGoods) {
	s.OrderGoods = data
}

func (s *CreateOrder) SetOrderGoodsStructure(data []*models.OrderGoodsStructure) {
	s.OrderGoodsStructure = data
}

func (s *CreateOrder) Process() (*models.Order, error) {
	engine := db.DB().Engine
	affected, err := engine.InsertOne(s.Order)
	if err != nil {
		return nil, util.NewError("订单保存错误")
	}
	log.Debugf("s.Order affected ", affected)
	s.OrderConsignee.Id = s.Order.Id
	_, err = engine.InsertOne(s.OrderConsignee)
	if err != nil {
		return nil, util.NewError("订单收货人保存错误")
	}
	s.OrderExt.Id = s.Order.Id
	_, err = engine.InsertOne(s.OrderExt)
	if err != nil {
		return nil, util.NewError("订单扩展信息保存错误")
	}
	for _, goods := range s.OrderGoods {
		goods.OrderId = s.Order.Id
		_, err = engine.Insert(goods)
		if err != nil {
			return nil, util.NewError("订单商品保存错误")
		}
	}
	for _, goods := range s.OrderGoodsStructure {
		goods.OrderId = s.Order.Id
		_, err = engine.Insert(goods)
		if err != nil {
			return nil, util.NewError("订单商品保存错误")
		}
	}

	return s.Order, nil
}
