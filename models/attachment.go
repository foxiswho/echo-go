package models

import (
	"time"
)

type Attachment struct {
	Id           int       `json:"id" xorm:"not null pk autoincr comment('附件ID') INT(10)"`
	Module       string    `json:"module" xorm:"comment('模块') index CHAR(32)"`
	Mark         string    `json:"mark" xorm:"comment('标记标志') index CHAR(60)"`
	TypeId       int       `json:"type_id" xorm:"not null default 0 comment('类别ID') INT(5)"`
	Name         string    `json:"name" xorm:"comment('保存的文件名称') CHAR(50)"`
	NameOriginal string    `json:"name_original" xorm:"comment('原文件名') VARCHAR(255)"`
	Path         string    `json:"path" xorm:"comment('文件路径') CHAR(200)"`
	Size         int       `json:"size" xorm:"not null default 0 comment('文件大小') INT(10)"`
	Ext          string    `json:"ext" xorm:"comment('文件后缀') CHAR(10)"`
	IsImage      int       `json:"is_image" xorm:"not null default 0 comment('是否图片1是0否') TINYINT(1)"`
	IsThumb      int       `json:"is_thumb" xorm:"not null default 0 comment('是否缩略图1是0否') TINYINT(1)"`
	Downloads    int       `json:"downloads" xorm:"not null default 0 comment('下载次数') INT(8)"`
	GmtCreate    time.Time `json:"gmt_create" xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间上传时间') TIMESTAMP"`
	Ip           string    `json:"ip" xorm:"comment('上传IP') CHAR(15)"`
	Status       int       `json:"status" xorm:"not null default 0 comment('状态99正常;') index TINYINT(2)"`
	Md5          string    `json:"md5" xorm:"comment('md5') index CHAR(32)"`
	Sha1         string    `json:"sha1" xorm:"comment('sha1') CHAR(40)"`
	FromId       int       `json:"from_id" xorm:"not null default 0 comment('所属ID') index INT(10)"`
	Aid          int       `json:"aid" xorm:"not null default 0 comment('后台管理员ID') index INT(10)"`
	Uid          int       `json:"uid" xorm:"not null default 0 comment('前台用户ID') index INT(10)"`
	IsShow       int       `json:"is_show" xorm:"not null default 1 comment('是否显示1是0否') index TINYINT(1)"`
	Http         string    `json:"http" xorm:"comment('图片http地址') VARCHAR(100)"`
}

//初始化
func NewAttachment() *Attachment {
	return new(Attachment)
}
