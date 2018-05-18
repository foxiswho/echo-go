package models

import (
	"time"
)

type Log struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	FromId     int       `json:"from_id" xorm:"not null default 0 comment('id') index INT(11)"`
	Aid        int       `json:"aid" xorm:"not null default 0 comment('管理员ID') index INT(11)"`
	Uid        int       `json:"uid" xorm:"not null default 0 comment('用户id') index INT(11)"`
	GmtCreate  time.Time `json:"gmt_create" xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	Mark       string    `json:"mark" xorm:"comment('标志自定义标志') CHAR(32)"`
	Data       string    `json:"data" xorm:"comment('其他内容') TEXT"`
	No         string    `json:"no" xorm:"comment('单号') index CHAR(50)"`
	TypeLogin  int       `json:"type_login" xorm:"not null default 0 comment('登录方式;302前台还是后台301') index INT(11)"`
	TypeClient int       `json:"type_client" xorm:"not null default 0 comment('登录客户端类别;321电脑;322安卓;323IOS;324手机网页;325其他') index INT(11)"`
	Ip         string    `json:"ip" xorm:"comment('IP') CHAR(20)"`
	Msg        string    `json:"msg" xorm:"comment('自定义说明') VARCHAR(255)"`
}

//初始化
func NewLog() *Log {
	return new(Log)
}
