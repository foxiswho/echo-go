package models

import (
	"time"
)

type GoodsCategory struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Name        string    `json:"name" xorm:"comment('名称') VARCHAR(50)"`
	Description string    `json:"description" xorm:"comment('介绍') TEXT"`
	ParentId    int       `json:"parent_id" xorm:"not null default 0 index INT(10)"`
	Sort        int       `json:"sort" xorm:"not null default 0 comment('排序') index INT(10)"`
	ArrParentId string    `json:"arr_parent_id" xorm:"comment('所有父栏目ID') VARCHAR(255)"`
	IsChild     int       `json:"is_child" xorm:"not null default 0 comment('是否有子栏目') TINYINT(1)"`
	ArrChildId  string    `json:"arr_child_id" xorm:"comment('所有子栏目ID') TEXT"`
	IsDel       int       `json:"is_del" xorm:"not null default 0 comment('是否删除1是0否') TINYINT(1)"`
	GmtCreate   time.Time `json:"gmt_create" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	GmtModified time.Time `json:"gmt_modified" xorm:"default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
}

//初始化
func NewGoodsCategory() *GoodsCategory {
	return new(GoodsCategory)
}
