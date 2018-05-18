package models

type UserGroupExt struct {
	GroupId int `json:"group_id" xorm:"not null pk autoincr comment('用户ID') INT(10)"`
}

//初始化
func NewUserGroupExt() *UserGroupExt {
	return new(UserGroupExt)
}
