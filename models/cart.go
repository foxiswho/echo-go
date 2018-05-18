package models

type Cart struct {
	Id          int   `json:"id" xorm:"not null pk autoincr INT(10)"`
	Uid         int   `json:"uid" xorm:"not null default 0 comment('用户ID') index INT(10)"`
	GoodsId     int   `json:"goods_id" xorm:"not null default 0 comment('商品ID') INT(10)"`
	ProductId   int   `json:"product_id" xorm:"not null default 0 comment('商品信息id') INT(10)"`
	Num         int64 `json:"num" xorm:"not null default 0 comment('数量') BIGINT(10)"`
	Price       int64 `json:"price" xorm:"not null default 0 comment('单价') BIGINT(20)"`
	Amount      int64 `json:"amount" xorm:"not null default 0 comment('合计总价') BIGINT(20)"`
	WarehouseId int   `json:"warehouse_id" xorm:"not null default 0 comment('仓库ID') INT(10)"`
	Sid         int   `json:"sid" xorm:"not null default 0 comment('供货商ID') INT(10)"`
	TypeId      int   `json:"type_id" xorm:"not null default 1 comment('类别:1普通') index INT(11)"`
}

//初始化
func NewCart() *Cart {
	return new(Cart)
}
