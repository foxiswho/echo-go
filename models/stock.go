package models

import (
	"time"
)

type Stock struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	WarehouseId  int       `json:"warehouse_id" xorm:"not null default 0 comment('仓库ID') index(warehouse_product_wms_uid) INT(11)"`
	ProductId    int       `json:"product_id" xorm:"not null default 0 comment('商品ID') index(warehouse_product_wms_uid) INT(10)"`
	Sid          int       `json:"sid" xorm:"not null default 0 comment('供应商ID') index(warehouse_product_wms_uid) INT(11)"`
	Number       string    `json:"number" xorm:"comment('商品编号') CHAR(100)"`
	Barcode      string    `json:"barcode" xorm:"comment('条形码') CHAR(32)"`
	Title        string    `json:"title" xorm:"comment('商品标题') VARCHAR(200)"`
	Num          int       `json:"num" xorm:"not null default 0 comment('数量(展示库存)') INT(11)"`
	IsUserLock   int       `json:"is_user_lock" xorm:"not null default 0 comment('是否有用户锁定库存') index TINYINT(1)"`
	TypeId       int       `json:"type_id" xorm:"not null default 0 comment('类别') INT(11)"`
	NumAvailable int       `json:"num_available" xorm:"not null default 0 comment('可用数量') INT(11)"`
	NumLocking   int       `json:"num_locking" xorm:"not null default 0 comment('锁定库存') INT(10)"`
	NumUserLock  int       `json:"num_user_lock" xorm:"not null default 0 comment('用户锁定数量') INT(11)"`
	NumWms       int       `json:"num_wms" xorm:"not null default 0 comment('wms库存=可用数量+用户锁定数量-锁定库存') INT(10)"`
	Mark         string    `json:"mark" xorm:"comment('标志') unique CHAR(32)"`
	GmtCreate    time.Time `json:"gmt_create" xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	GmtModified  time.Time `json:"gmt_modified" xorm:"default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
}

//初始化
func NewStock() *Stock {
	return new(Stock)
}
