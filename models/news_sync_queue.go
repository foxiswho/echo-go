package models

import (
	"time"
)

type NewsSyncQueue struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	NewsId      int       `json:"news_id" xorm:"not null default 0 comment('本站博客id') INT(11)"`
	TypeId      int       `json:"type_id" xorm:"not null default 0 comment('类型') INT(11)"`
	Status      int       `json:"status" xorm:"not null default 0 comment('状态：0:待运行 10:失败 99:成功') TINYINT(3)"`
	GmtModified time.Time `json:"gmt_modified" xorm:"default 'CURRENT_TIMESTAMP' comment('最后一次更新时间') TIMESTAMP"`
	GmtCreate   time.Time `json:"gmt_create" xorm:"default 'CURRENT_TIMESTAMP' comment('插入时间') TIMESTAMP"`
	Msg         string    `json:"msg" xorm:"comment('内容') VARCHAR(255)"`
	MapId       int       `json:"map_id" xorm:"not null default 0 comment('同步ID') INT(11)"`
}

//初始化
func NewNewsSyncQueue() *NewsSyncQueue {
	return new(NewsSyncQueue)
}
