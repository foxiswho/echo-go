package models

import (
	"time"
)

type Tag struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name      string    `json:"name" xorm:"comment('名称') CHAR(50)"`
	GmtCreate time.Time `json:"gmt_create" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
}

//初始化
func NewTag() *Tag {
	return new(Tag)
}
