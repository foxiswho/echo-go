package models

import (
	"time"
)

type NewsTag struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name      string    `json:"name" xorm:"comment('名称') CHAR(100)"`
	GmtCreate time.Time `json:"gmt_create" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	Aid       int       `json:"aid" xorm:"not null default 0 comment('管理员ID') INT(11)"`
	NewsId    int       `json:"news_id" xorm:"not null default 0 comment('文章ID') INT(11)"`
}

//初始化
func NewNewsTag() *NewsTag {
	return new(NewsTag)
}
