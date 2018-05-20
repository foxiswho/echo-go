package order_service

import (
	"github.com/foxiswho/echo-go/dao/order_dao"
	"github.com/foxiswho/echo-go/models"
	"github.com/foxiswho/echo-go/util"
)

//创建订单前格式化
type CreateOrderFormat struct {
	OrderGoodsData []*models.OrderGoodsData //数据集合
}

func (s *CreateOrderFormat) Process() (*models.Order, error) {
	if s.OrderGoodsData == nil {
		return nil, util.NewError("商品数据不能为空")
	}
	if len(s.OrderGoodsData) == 0 {
		return nil, util.NewError("商品数据不能为空")
	}

	order := order_dao.NewCreateOrder()

	order.OrderGoods = make([]*models.OrderGoods, 0)
	for _, goods := range s.OrderGoodsData {
		order_goods := models.NewOrderGoods()
		order_goods.Num = goods.Num
		order_goods.Number = goods.Number
		order_goods.NumUnit = goods.NumUnit
		order_goods.NumTotal = order_goods.NumUnit * order_goods.Num
		order_goods.Sid = goods.Sid
		order_goods.ProductId = goods.ProductId
		order_goods.WarehouseId = goods.WarehouseId
		order_goods.GoodsId = goods.Goods.Id
		order_goods.MarkId = goods.MarkId
		order_goods.Price = goods.Price
		order_goods.PriceShop = goods.PriceShop
		order.OrderGoods = append(order.OrderGoods, order_goods)
	}

	return order.Process()
}
