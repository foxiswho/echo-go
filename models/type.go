package models

import (
	"time"
)

type Type struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name      string    `json:"name" xorm:"comment('名称') CHAR(100)"`
	NameEn    string    `json:"name_en" xorm:"comment('名称') CHAR(100)"`
	Code      string    `json:"code" xorm:"comment('代码') CHAR(32)"`
	Mark      string    `json:"mark" xorm:"comment('标志') index CHAR(32)"`
	TypeId    int       `json:"type_id" xorm:"not null default 0 comment('所属类别ID') index INT(11)"`
	ParentId  int       `json:"parent_id" xorm:"not null default 0 comment('上级ID、属于/上级ID') index INT(11)"`
	Value     int       `json:"value" xorm:"not null default 0 comment('值') INT(10)"`
	Content   string    `json:"content" xorm:"comment('内容') VARCHAR(255)"`
	IsDel     int       `json:"is_del" xorm:"not null default 0 comment('是否删除0否1是') index INT(11)"`
	Sort      int       `json:"sort" xorm:"not null default 0 comment('排序') index INT(11)"`
	Remark    string    `json:"remark" xorm:"comment('备注') VARCHAR(255)"`
	GmtCreate time.Time `json:"gmt_create" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	Aid       int       `json:"aid" xorm:"not null default 0 comment('添加人') INT(11)"`
	Module    string    `json:"module" xorm:"comment('模块') CHAR(50)"`
	Setting   string    `json:"setting" xorm:"comment('扩展参数') VARCHAR(255)"`
	IsDefault int       `json:"is_default" xorm:"not null default 0 comment('是否默认') TINYINT(1)"`
	IsChild   int       `json:"is_child" xorm:"not null default 0 comment('是否有子类1是0否') TINYINT(1)"`
	IsSystem  int       `json:"is_system" xorm:"not null default 0 comment('系统参数禁止修改') TINYINT(1)"`
	IsShow    int       `json:"is_show" xorm:"not null default 0 comment('是否显示在配置页面上') TINYINT(1)"`
}

//初始化
func NewType() *Type {
	return new(Type)
}
