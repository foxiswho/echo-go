package models

import (
	"time"
)

type GoodsPrice struct {
	Id               int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	PriceMarket      int64     `json:"price_market" xorm:"not null default 0 comment('市场价') BIGINT(12)"`
	PriceShop        int64     `json:"price_shop" xorm:"not null default 0 comment('商城价') BIGINT(12)"`
	IsPromote        int       `json:"is_promote" xorm:"not null default 0 comment('是否促销1是0否') TINYINT(1)"`
	PromotePrice     int       `json:"promote_price" xorm:"not null default 0 comment('促销价格') INT(12)"`
	PromoteStartDate time.Time `json:"promote_start_date" xorm:"comment('促销开始日期') TIMESTAMP"`
	PromoteEndDate   time.Time `json:"promote_end_date" xorm:"comment('促销结束日期') TIMESTAMP"`
	IsFreeShipping   int       `json:"is_free_shipping" xorm:"not null default 0 comment('是否包邮1是0否') TINYINT(1)"`
	StartDate        time.Time `json:"start_date" xorm:"comment('开始时间') TIMESTAMP"`
	EndDate          time.Time `json:"end_date" xorm:"comment('结束时间') TIMESTAMP"`
	MinFreeShipping  int       `json:"min_free_shipping" xorm:"not null default 1 comment('最小包邮数量') INT(10)"`
	NumMax           string    `json:"num_max" xorm:"not null default '9999' comment('最大可一次购买数量') VARCHAR(255)"`
	NumLeast         int       `json:"num_least" xorm:"not null default 1 comment('最少购买数量') INT(11)"`
	IsFreeTax        int       `json:"is_free_tax" xorm:"not null default 0 comment('是否包税使用包税价格') TINYINT(1)"`
	IsGroupPrice     string    `json:"is_group_price" xorm:"default 1 comment('是否使用用户组价格') DECIMAL(10)"`
	TaxPrice         int64     `json:"tax_price" xorm:"not null default 0 comment('包税价格') BIGINT(12)"`
}

//初始化
func NewGoodsPrice() *GoodsPrice {
	return new(GoodsPrice)
}
