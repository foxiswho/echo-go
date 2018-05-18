package models

type AdminRoleAccess struct {
	Id        int `json:"id" xorm:"not null pk autoincr INT(11)"`
	Aid       int `json:"aid" xorm:"not null default 0 comment('管理员ID') unique(aid_role_id) INT(11)"`
	RoleId    int `json:"role_id" xorm:"not null default 0 comment('角色ID') unique(aid_role_id) INT(11)"`
	IsDefault int `json:"is_default" xorm:"not null default 0 comment('是否默认') TINYINT(1)"`
}

//初始化
func NewAdminRoleAccess() *AdminRoleAccess {
	return new(AdminRoleAccess)
}
