package models

import (
	"time"
)

type Connect struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	TypeId    int       `json:"type_id" xorm:"not null default 0 comment('类别id') index INT(11)"`
	Uid       int       `json:"uid" xorm:"not null default 0 comment('用户id') index INT(11)"`
	OpenId    string    `json:"open_id" xorm:"comment('对应唯一开放id') index CHAR(80)"`
	Token     string    `json:"token" xorm:"comment('开放密钥') VARCHAR(80)"`
	Type      int       `json:"type" xorm:"not null default 1 comment('登录类型1腾讯QQ2新浪微博') INT(11)"`
	TypeLogin int       `json:"type_login" xorm:"not null default 0 comment('登录模块;302前台还是后台301') INT(11)"`
	GmtCreate time.Time `json:"gmt_create" xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	Extend    string    `json:"extend" xorm:"default '' comment('扩展参数') VARCHAR(5000)"`
}

//初始化
func NewConnect() *Connect {
	return new(Connect)
}
