package models

type Area struct {
	Id              int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name            string `json:"name" xorm:"default '' comment('名称') CHAR(50)"`
	NameEn          string `json:"name_en" xorm:"default '' comment('英文名称') VARCHAR(100)"`
	ParentId        int    `json:"parent_id" xorm:"default 0 comment('上级栏目ID') index INT(11)"`
	Type            int    `json:"type" xorm:"default 0 comment('类别;0默认;') TINYINT(4)"`
	NameTraditional string `json:"name_traditional" xorm:"default '' comment('繁体名称') VARCHAR(50)"`
	Sort            int    `json:"sort" xorm:"default 0 comment('排序') INT(11)"`
}

//初始化
func NewArea() *Area {
	return new(Area)
}
