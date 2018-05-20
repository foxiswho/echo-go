package order_service

import (
	"github.com/foxiswho/echo-go/models"
	"github.com/foxiswho/echo-go/util"
	"github.com/foxiswho/echo-go/service"
)

//创建订单，数据来源于购物车
type CreateOrderFormatFromCart struct {
	Cart []*models.Cart
}

func NewCreateOrderFormatFromCart() *CreateOrderFormatFromCart {
	return new(CreateOrderFormatFromCart)
}

func (s *CreateOrderFormatFromCart) SetCart(cart []*models.Cart) {
	s.Cart = cart
}

func (s *CreateOrderFormatFromCart) Process() ([]*models.OrderGoodsData, error) {
	if s.Cart == nil {
		return nil, util.NewError("购物车商品数据不能为空")
	}
	if len(s.Cart) == 0 {
		return nil, util.NewError("购物车商品数据不能为空")
	}
	order_goods_data_all := make([]*models.OrderGoodsData, 0)
	for _, val := range s.Cart {
		order_goods_data := models.NewOrderGoodsData()
		goods, err := service.NewGoodsService().GetById(val.GoodsId)
		if err != nil {
			return nil, err
		}
		goods_price, err := service.NewGoodsPriceService().GetById(val.GoodsId)
		if err != nil {
			return nil, err
		}
		order_goods_data.Goods = goods
		order_goods_data.GoodsPrice = goods_price
		order_goods_data.Num = val.Num
		order_goods_data_all = append(order_goods_data_all, order_goods_data)
	}
	return order_goods_data_all, nil
}
