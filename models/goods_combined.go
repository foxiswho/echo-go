package models

type GoodsCombined struct {
	Id        int   `json:"id" xorm:"not null pk autoincr INT(11)"`
	GoodsId   int   `json:"goods_id" xorm:"not null default 0 INT(11)"`
	ProductId int   `json:"product_id" xorm:"not null default 0 comment('产品ID') INT(11)"`
	Aid       int   `json:"aid" xorm:"not null default 0 comment('添加人') INT(11)"`
	PriceShop int64 `json:"price_shop" xorm:"not null default 0 comment('组合商品价格') BIGINT(20)"`
	Sort      int   `json:"sort" xorm:"not null default 0 comment('排序') INT(11)"`
	NumLeast  int   `json:"num_least" xorm:"not null default 1 comment('最少购买数量') INT(11)"`
	ParentId  int   `json:"parent_id" xorm:"not null default 0 comment('顶级商品ID') INT(11)"`
}

//初始化
func NewGoodsCombined() *GoodsCombined {
	return new(GoodsCombined)
}
