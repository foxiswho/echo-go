package models

//订单商品数据集合
type OrderGoodsData struct {
	*Goods
	*GoodsPrice
	Num   int //	数量
	Price int64 //	价格
}

//初始化
func NewOrderGoodsData() *OrderGoodsData {
	return new(OrderGoodsData)
}
