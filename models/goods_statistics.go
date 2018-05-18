package models

type GoodsStatistics struct {
	Id      int   `json:"id" xorm:"not null pk INT(11)"`
	Saless  int64 `json:"saless" xorm:"not null default 0 comment('销量') BIGINT(20)"`
	Reading int64 `json:"reading" xorm:"not null default 0 comment('访问数') BIGINT(20)"`
}

//初始化
func NewGoodsStatistics() *GoodsStatistics {
	return new(GoodsStatistics)
}
