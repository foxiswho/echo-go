package models

type AreaExt struct {
	Id              int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	AreaId          int    `json:"area_id" xorm:"default 0 comment('ID') index(id) INT(11)"`
	Name            string `json:"name" xorm:"default '' comment('名称') CHAR(50)"`
	NameEn          string `json:"name_en" xorm:"default '' comment('英文名称') VARCHAR(100)"`
	ParentId        int    `json:"parent_id" xorm:"default 0 comment('上级栏目ID') index(id) INT(11)"`
	Type            int    `json:"type" xorm:"default 0 comment('类别;0默认;1又名;2;3属于;11已合并到;12已更名为') TINYINT(4)"`
	NameTraditional string `json:"name_traditional" xorm:"default '' comment('繁体名称') VARCHAR(50)"`
	Sort            int    `json:"sort" xorm:"default 0 comment('排序') INT(11)"`
	TypeName        string `json:"type_name" xorm:"default '' comment('类别名称') VARCHAR(50)"`
	OtherName       string `json:"other_name" xorm:"default '' comment('根据类别名称填写') VARCHAR(50)"`
}

//初始化
func NewAreaExt() *AreaExt {
	return new(AreaExt)
}
