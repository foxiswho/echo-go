package models

type OrderConsignee struct {
	Id           int    `json:"id" xorm:"not null pk autoincr comment('订单id') INT(11)"`
	ConsigneeId  int    `json:"consignee_id" xorm:"not null default 0 comment('收货人ID') INT(10)"`
	Consignee    string `json:"consignee" xorm:"comment('收货人') VARCHAR(50)"`
	Mobile       string `json:"mobile" xorm:"comment('手机号') CHAR(11)"`
	Country      int    `json:"country" xorm:"not null default 1 comment('国家') INT(11)"`
	Province     int    `json:"province" xorm:"not null default 0 comment('省') INT(11)"`
	City         int    `json:"city" xorm:"not null default 0 comment('市') INT(11)"`
	District     int    `json:"district" xorm:"not null default 0 comment('区') INT(11)"`
	Address      string `json:"address" xorm:"comment('地址') VARCHAR(255)"`
	AddressEn    string `json:"address_en" xorm:"comment('地址(英文)') VARCHAR(255)"`
	IdCard       string `json:"id_card" xorm:"comment('身份证号') CHAR(19)"`
	IdCardFront  string `json:"id_card_front" xorm:"comment('身份证正面') VARCHAR(255)"`
	IdCardBack   string `json:"id_card_back" xorm:"comment('身份证反面') VARCHAR(255)"`
	ProvinceName string `json:"province_name" xorm:"comment('省') CHAR(30)"`
	CityName     string `json:"city_name" xorm:"comment('市') CHAR(50)"`
	DistrictName string `json:"district_name" xorm:"comment('区') CHAR(50)"`
	AddressAll   string `json:"address_all" xorm:"comment('地址') VARCHAR(255)"`
}

//初始化
func NewOrderConsignee() *OrderConsignee {
	return new(OrderConsignee)
}
