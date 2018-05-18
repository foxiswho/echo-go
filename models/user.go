package models

import (
	"time"
)

type User struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Mobile      string    `json:"mobile" xorm:"index CHAR(11)"`
	Username    string    `json:"username" xorm:"comment('用户名') index CHAR(30)"`
	Mail        string    `json:"mail" xorm:"comment('邮箱') index CHAR(32)"`
	Password    string    `json:"password" xorm:"comment('密码') CHAR(32)"`
	Salt        string    `json:"salt" xorm:"comment('干扰码') CHAR(6)"`
	RegIp       string    `json:"reg_ip" xorm:"comment('注册IP') CHAR(15)"`
	RegTime     time.Time `json:"reg_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('注册时间') TIMESTAMP"`
	IsDel       int       `json:"is_del" xorm:"not null default 0 comment('状态0正常1删除') index TINYINT(1)"`
	GroupId     int       `json:"group_id" xorm:"not null default 410 comment('用户组ID') index INT(11)"`
	TrueName    string    `json:"true_name" xorm:"comment('真实姓名') VARCHAR(32)"`
	Name        string    `json:"name" xorm:"comment('店铺名称') VARCHAR(100)"`
	GmtCreate   time.Time `json:"gmt_create" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	GmtModified time.Time `json:"gmt_modified" xorm:"default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
}

//初始化
func NewUser() *User {
	return new(User)
}
