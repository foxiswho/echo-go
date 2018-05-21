package models

type OrderCollectDate struct {
	Order               *Order
	OrderExt            *OrderExt
	Consignee           *OrderConsignee
	OrderGoods          []*OrderGoods
	OrderGoodsStructure []*OrderGoodsStructure
}
