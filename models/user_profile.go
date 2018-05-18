package models

type UserProfile struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Sex         int    `json:"sex" xorm:"not null default 1 comment('性别1男2女3中性0保密') TINYINT(1)"`
	Job         string `json:"job" xorm:"comment('担任职务') VARCHAR(50)"`
	Qq          string `json:"qq" xorm:"VARCHAR(20)"`
	Phone       string `json:"phone" xorm:"comment('电话') VARCHAR(20)"`
	County      int    `json:"county" xorm:"not null default 1 comment('国家') INT(11)"`
	Province    int    `json:"province" xorm:"not null default 0 comment('省') INT(11)"`
	City        int    `json:"city" xorm:"not null default 0 comment('市') INT(11)"`
	District    int    `json:"district" xorm:"not null default 0 comment('区') INT(11)"`
	Address     string `json:"address" xorm:"comment('地址') VARCHAR(255)"`
	Wechat      string `json:"wechat" xorm:"comment('微信') VARCHAR(20)"`
	RemarkAdmin string `json:"remark_admin" xorm:"comment('客服备注') TEXT"`
}

//初始化
func NewUserProfile() *UserProfile {
	return new(UserProfile)
}
