package models

import (
	"time"
)

type UserStatus struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	RegIp          string    `json:"reg_ip" xorm:"comment('注册IP') CHAR(15)"`
	RegTime        time.Time `json:"reg_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('注册时间') TIMESTAMP"`
	RegType        int       `json:"reg_type" xorm:"not null default 0 comment('注册方式') INT(11)"`
	RegAppId       int       `json:"reg_app_id" xorm:"not null default 1 comment('注册来源') INT(11)"`
	LastLoginIp    string    `json:"last_login_ip" xorm:"comment('最后登录IP') CHAR(15)"`
	LastLoginTime  time.Time `json:"last_login_time" xorm:"comment('最后登录时间') TIMESTAMP"`
	LastLoginAppId int       `json:"last_login_app_id" xorm:"not null default 0 comment('最后登录app_id') INT(11)"`
	Login          int       `json:"login" xorm:"not null default 0 comment('登录次数') SMALLINT(5)"`
	IsMobile       int       `json:"is_mobile" xorm:"not null default 0 comment('手机号是否已验证1已验证0未验证') TINYINT(1)"`
	IsEmail        int       `json:"is_email" xorm:"not null default 0 comment('邮箱是否已验证1已验证0未验证') TINYINT(1)"`
	AidAdd         int       `json:"aid_add" xorm:"not null default 0 comment('客服AID') INT(11)"`
}

//初始化
func NewUserStatus() *UserStatus {
	return new(UserStatus)
}
