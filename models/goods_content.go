package models

type GoodsContent struct {
	Id             int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	SeoTitle       string `json:"seo_title" xorm:"comment('seo标题') VARCHAR(50)"`
	SeoDescription string `json:"seo_description" xorm:"comment('seo描述') VARCHAR(200)"`
	SeoKeyword     string `json:"seo_keyword" xorm:"comment('seo关键字') VARCHAR(50)"`
	Content        string `json:"content" xorm:"comment('内容') TEXT"`
	Remark         string `json:"remark" xorm:"comment('备注紧供自己查看') VARCHAR(255)"`
	TitleOther     string `json:"title_other" xorm:"comment('其他名称') VARCHAR(5000)"`
}

//初始化
func NewGoodsContent() *GoodsContent {
	return new(GoodsContent)
}
