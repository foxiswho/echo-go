package models

type OrderExt struct {
	Id                   int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	CostAmountPayment    int64  `json:"cost_amount_payment" xorm:"not null default 0 comment('支付金额成本') BIGINT(20)"`
	CostAmountGoods      int64  `json:"cost_amount_goods" xorm:"not null default 0 comment('商品金额成本') BIGINT(20)"`
	DeclareTax           int    `json:"declare_tax" xorm:"not null default 0 comment('申报税') INT(11)"`
	DeclareVatFee        int    `json:"declare_vat_fee" xorm:"not null default 0 comment('申报增值税') INT(11)"`
	DeclareSalesFee      int    `json:"declare_sales_fee" xorm:"not null default 0 comment('申报消费税') INT(11)"`
	DeclareAmountFreight int    `json:"declare_amount_freight" xorm:"not null default 0 comment('申报运费') INT(11)"`
	DeclareAmountPayment int    `json:"declare_amount_payment" xorm:"not null default 0 comment('申报运费') INT(11)"`
	IsWmsSales           int    `json:"is_wms_sales" xorm:"not null default 0 comment('是否向WMS通信生成销售单   0未通信 1已通信') TINYINT(4)"`
	IsWmsSendOut         int    `json:"is_wms_send_out" xorm:"not null default 0 comment('是否向WMS发送并生成出货单  0未完成 1已完成') TINYINT(4)"`
	OrderAmountDeclare   int64  `json:"order_amount_declare" xorm:"not null default 0 comment('订单总金额') BIGINT(20)"`
	PaymentAmountDeclare int64  `json:"payment_amount_declare" xorm:"not null default 0 comment('支付总金额') BIGINT(20)"`
	GoodsAmountDeclare   int64  `json:"goods_amount_declare" xorm:"not null default 0 comment('商品小计') BIGINT(20)"`
	BillingCountry       int    `json:"billing_country" xorm:"not null default 1 comment('账单国家') INT(11)"`
	BillingProvince      int    `json:"billing_province" xorm:"not null default 0 comment('账单省') INT(11)"`
	BillingCity          int    `json:"billing_city" xorm:"not null default 0 comment('账单市') INT(11)"`
	BillingDistrict      int    `json:"billing_district" xorm:"not null default 0 comment('账单区') INT(11)"`
	BillingAddress       string `json:"billing_address" xorm:"not null default '' comment('账单地址') VARCHAR(255)"`
	BillingMobile        string `json:"billing_mobile" xorm:"not null default '' comment('手机号') CHAR(11)"`
	BillingConsignee     string `json:"billing_consignee" xorm:"not null default '' comment('账单收货人') VARCHAR(255)"`
	BillingMail          string `json:"billing_mail" xorm:"not null default '' comment('账单邮箱') VARCHAR(255)"`
	BillingAddressEn     string `json:"billing_address_en" xorm:"comment('账单地址(英文)') VARCHAR(255)"`
	BillingZipCode       string `json:"billing_zip_code" xorm:"comment('账单邮编') VARCHAR(10)"`
	BillingTaxNo         string `json:"billing_tax_no" xorm:"comment('税号') VARCHAR(255)"`
	PackingId            int    `json:"packing_id" xorm:"default 0 comment('包装ID') INT(10)"`
}

//初始化
func NewOrderExt() *OrderExt {
	return new(OrderExt)
}
