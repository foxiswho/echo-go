package models

import (
	"time"
)

type AdminStatus struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	LoginTime   time.Time `json:"login_time" xorm:"comment('登录时间') TIMESTAMP"`
	LoginIp     string    `json:"login_ip" xorm:"comment('IP') CHAR(20)"`
	Login       int       `json:"login" xorm:"not null default 0 comment('登录次数') INT(11)"`
	AidAdd      int       `json:"aid_add" xorm:"not null default 0 comment('添加人') INT(11)"`
	AidUpdate   int       `json:"aid_update" xorm:"not null default 0 comment('更新人') INT(11)"`
	GmtModified time.Time `json:"gmt_modified" xorm:"comment('更新时间') TIMESTAMP"`
	Remark      string    `json:"remark" xorm:"comment('备注') VARCHAR(255)"`
}

//初始化
func NewAdminStatus() *AdminStatus {
	return new(AdminStatus)
}
