package models

type OrderGoodsStructure struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	OrderId     int    `json:"order_id" xorm:"not null default 0 comment('订单ID') index INT(10)"`
	GoodsId     int    `json:"goods_id" xorm:"not null default 0 comment('商品ID') INT(10)"`
	ProductId   int    `json:"product_id" xorm:"not null default 0 comment('商品信息id') INT(10)"`
	Title       string `json:"title" xorm:"comment('商品名称') VARCHAR(200)"`
	Num         int    `json:"num" xorm:"not null default 0 comment('数量') INT(10)"`
	Number      string `json:"number" xorm:"comment('商品编号') CHAR(100)"`
	Price       int64  `json:"price" xorm:"not null default 0 comment('单价') BIGINT(20)"`
	NumUnit     int    `json:"num_unit" xorm:"not null default 1 comment('每个单位内多少个，每盒几罐') INT(11)"`
	NumTotal    int    `json:"num_total" xorm:"not null default 0 comment('总数量 = 罐数x页面数量') INT(11)"`
	Amount      int64  `json:"amount" xorm:"not null default 0 comment('合计总价') BIGINT(20)"`
	Freight     int64  `json:"freight" xorm:"not null default 0 comment('运费') BIGINT(20)"`
	WarehouseId int    `json:"warehouse_id" xorm:"not null default 0 comment('仓库ID') INT(10)"`
	Sid         int    `json:"sid" xorm:"not null default 1 comment('商家ID') INT(10)"`
	SalesFee    int64  `json:"sales_fee" xorm:"not null default 0 comment('消费税费') BIGINT(11)"`
	VatFee      int64  `json:"vat_fee" xorm:"not null default 0 comment('增值税费') BIGINT(10)"`
	PriceTax    int64  `json:"price_tax" xorm:"not null default 0 comment('总税费') BIGINT(10)"`
	Remark      string `json:"remark" xorm:"comment('备注') TEXT"`
	PriceShop   int64  `json:"price_shop" xorm:"default 0 comment('商城价') BIGINT(12)"`
	CostPrice   int64  `json:"cost_price" xorm:"not null default 0 comment('成本单价') BIGINT(11)"`
	CostAmount  int64  `json:"cost_amount" xorm:"not null default 0 comment('成本金额') BIGINT(20)"`
	ParentId    int    `json:"parent_id" xorm:"not null default 0 comment('所属组合商品ID') INT(11)"`
}

//初始化
func NewOrderGoodsStructure() *OrderGoodsStructure {
	return new(OrderGoodsStructure)
}
