package models

import (
	"time"
)

type Admin struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Username     string    `json:"username" xorm:"comment('用户名') index CHAR(30)"`
	Password     string    `json:"password" xorm:"comment('密码') CHAR(32)"`
	Mail         string    `json:"mail" xorm:"comment('邮箱') VARCHAR(80)"`
	Salt         string    `json:"salt" xorm:"comment('干扰码') VARCHAR(10)"`
	GmtCreate    time.Time `json:"gmt_create" xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	GmtModified  time.Time `json:"gmt_modified" xorm:"default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	Ip           string    `json:"ip" xorm:"comment('添加IP') CHAR(15)"`
	JobNo        string    `json:"job_no" xorm:"comment('工号') VARCHAR(15)"`
	NickName     string    `json:"nick_name" xorm:"comment('昵称') VARCHAR(50)"`
	TrueName     string    `json:"true_name" xorm:"comment('真实姓名') VARCHAR(50)"`
	Qq           string    `json:"qq" xorm:"comment('qq') VARCHAR(50)"`
	Phone        string    `json:"phone" xorm:"comment('电话') VARCHAR(50)"`
	Mobile       string    `json:"mobile" xorm:"comment('手机') VARCHAR(20)"`
	Name         string    `json:"name" xorm:"comment('显示名称') VARCHAR(255)"`
	IsDel        int       `json:"is_del" xorm:"not null default 0 comment('删除0否1是') index TINYINT(1)"`
	DepartmentId int       `json:"department_id" xorm:"not null default 0 comment('部门id') INT(11)"`
	TeamId       int       `json:"team_id" xorm:"not null comment('团队ID') INT(11)"`
	MasterId     int       `json:"master_id" xorm:"not null comment('师傅id') INT(11)"`
	LeaderId     int       `json:"leader_id" xorm:"not null comment('领导id') INT(11)"`
	PostId       int       `json:"post_id" xorm:"not null comment('职务id') INT(11)"`
	RoleId       int       `json:"role_id" xorm:"not null comment('角色id(主)') index INT(11)"`
}

//初始化
func NewAdmin() *Admin {
	return new(Admin)
}
