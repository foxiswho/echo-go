package models

type NewsStatistics struct {
	Id             int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	NewsId         int    `json:"news_id" xorm:"not null default 0 comment('文章ID') index INT(11)"`
	Comment        int    `json:"comment" xorm:"not null default 0 comment('评论人数') INT(11)"`
	Read           int    `json:"read" xorm:"not null default 0 comment('阅读人数') INT(11)"`
	SeoTitle       string `json:"seo_title" xorm:"comment('SEO标题') VARCHAR(255)"`
	SeoDescription string `json:"seo_description" xorm:"comment('SEO摘要') VARCHAR(255)"`
	SeoKeyword     string `json:"seo_keyword" xorm:"comment('SEO关键词') VARCHAR(255)"`
}

//初始化
func NewNewsStatistics() *NewsStatistics {
	return new(NewsStatistics)
}
