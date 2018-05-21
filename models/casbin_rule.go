package models

type CasbinRule struct {
	PType string `json:"p_type" xorm:"index VARCHAR(100)"`
	V0    string `json:"v0" xorm:"index VARCHAR(100)"`
	V1    string `json:"v1" xorm:"index VARCHAR(100)"`
	V2    string `json:"v2" xorm:"index VARCHAR(100)"`
	V3    string `json:"v3" xorm:"index VARCHAR(100)"`
	V4    string `json:"v4" xorm:"index VARCHAR(100)"`
	V5    string `json:"v5" xorm:"index VARCHAR(100)"`
	Id    int    `json:"id" xorm:"not null pk autoincr INT(1)"`
}

//初始化
func NewCasbinRule() *CasbinRule {
	return new(CasbinRule)
}
