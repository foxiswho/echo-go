package models

import (
	"time"
)

type Session struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Uid         int       `json:"uid" xorm:"not null default 0 comment('用户UID') index(uid) INT(11)"`
	Ip          string    `json:"ip" xorm:"comment('IP') CHAR(15)"`
	ErrorCount  int       `json:"error_count" xorm:"not null default 0 comment('密码输入错误次数') TINYINT(1)"`
	AppId       int       `json:"app_id" xorm:"not null default 0 comment('登录应用') INT(11)"`
	Md5         string    `json:"md5" xorm:"comment('md5') CHAR(32)"`
	TypeLogin   int       `json:"type_login" xorm:"not null default 0 comment('登录方式;302前台还是后台301') index(uid) INT(11)"`
	TypeClient  int       `json:"type_client" xorm:"not null default 0 comment('登录客户端类别;321电脑;322安卓;323IOS;324手机网页;325其他') index(uid) INT(11)"`
	GmtCreate   time.Time `json:"gmt_create" xorm:"default 'CURRENT_TIMESTAMP' comment('登录时间') TIMESTAMP"`
	GmtModified time.Time `json:"gmt_modified" xorm:"default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
}

//初始化
func NewSession() *Session {
	return new(Session)
}
