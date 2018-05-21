package models

type OrderCollectDate struct {
	Order               *Order
	OrderExt            *OrderExt
	OrderConsignee      *OrderConsignee
	OrderGoods          []*OrderGoods
	OrderGoodsStructure []*OrderGoodsStructure
}

func NewOrderCollectDate() *OrderCollectDate {
	return new(OrderCollectDate)
}
