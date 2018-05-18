package models

type AdminMenu struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name     string `json:"name" xorm:"comment('名称') CHAR(100)"`
	ParentId int    `json:"parent_id" xorm:"not null default 0 comment('上级菜单') index INT(11)"`
	S        string `json:"s" xorm:"comment('模块/控制器/动作') index CHAR(60)"`
	Data     string `json:"data" xorm:"comment('其他参数') CHAR(100)"`
	Sort     int    `json:"sort" xorm:"not null default 0 comment('排序') index INT(11)"`
	Remark   string `json:"remark" xorm:"comment('备注') VARCHAR(255)"`
	Type     string `json:"type" xorm:"not null default 'url' comment('类别url菜单function独立功能user用户独有') CHAR(32)"`
	Level    int    `json:"level" xorm:"not null default 0 comment('级别') TINYINT(3)"`
	Level1Id int    `json:"level1_id" xorm:"not null default 0 comment('1级栏目ID') INT(11)"`
	Md5      string `json:"md5" xorm:"comment('s的md5值') CHAR(32)"`
	IsShow   int    `json:"is_show" xorm:"not null default 1 comment('显示隐藏;1显示;0隐藏') TINYINT(1)"`
	IsUnique int    `json:"is_unique" xorm:"not null default 0 comment('用户独有此功能1是0否') TINYINT(1)"`
}

//初始化
func NewAdminMenu() *AdminMenu {
	return new(AdminMenu)
}
