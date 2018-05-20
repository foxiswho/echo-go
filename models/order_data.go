package models

type OrderDate struct {
	Order               *Order
	OrderExt            *OrderExt
	Consignee           *OrderConsignee
	OrderGoods          []*OrderGoods
	OrderGoodsStructure []*OrderGoodsStructure
}
