package models

import (
	"time"
)

type NewsSyncMapping struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	NewsId      int       `json:"news_id" xorm:"not null default 0 comment('本站blog的id') INT(11)"`
	TypeId      int       `json:"type_id" xorm:"not null default 0 comment('类别id') INT(11)"`
	ToId        string    `json:"to_id" xorm:"comment('csdn的id') VARCHAR(64)"`
	GmtModified time.Time `json:"gmt_modified" xorm:"default 'CURRENT_TIMESTAMP' comment('最后一次更新时间') TIMESTAMP"`
	GmtCreate   time.Time `json:"gmt_create" xorm:"default 'CURRENT_TIMESTAMP' comment('插入时间') TIMESTAMP"`
	Mark        string    `json:"mark" xorm:"comment('标志') CHAR(32)"`
	IsSync      int       `json:"is_sync" xorm:"not null default 0 comment('是否同步过') TINYINT(1)"`
	Extend      string    `json:"extend" xorm:"comment('扩展参数') VARCHAR(5000)"`
}

//初始化
func NewNewsSyncMapping() *NewsSyncMapping {
	return new(NewsSyncMapping)
}
