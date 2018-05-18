package models

type AdminRolePriv struct {
	Id     int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	RoleId int    `json:"role_id" xorm:"not null default 0 comment('角色ID') index index(role_id_2) SMALLINT(3)"`
	S      string `json:"s" xorm:"comment('模块/控制器/动作') index(role_id_2) CHAR(100)"`
	Data   string `json:"data" xorm:"comment('其他参数') CHAR(50)"`
	Aid    int    `json:"aid" xorm:"not null default 0 comment('管理员ID') INT(10)"`
	MenuId int    `json:"menu_id" xorm:"not null default 0 comment('菜单ID') INT(10)"`
	Type   string `json:"type" xorm:"not null default 'url' comment('类别url菜单function独立功能user用户独有') CHAR(32)"`
}

//初始化
func NewAdminRolePriv() *AdminRolePriv {
	return new(AdminRolePriv)
}
