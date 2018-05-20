package models

import (
	"time"
)

type Order struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	OrderNo        string    `json:"order_no" xorm:"comment('销售订单号') index CHAR(32)"`
	OrderSn        string    `json:"order_sn" xorm:"comment('单号淘宝苏宁等等') CHAR(32)"`
	Uid            int       `json:"uid" xorm:"not null default 0 comment('用户ID') index INT(10)"`
	OrderStatus    int       `json:"order_status" xorm:"not null default 0 comment('订单状态(DEFAULT用户未点发货,WAIT_CHECK等待审核,NO_CHECK审核不通过,WAIT_SEND等等发货,SEND卖家已发货,RECEIPT已收货,DROP交易关闭,SUCCESS订单交易成功,CANCEL交易取消,WAIT_CUSTOMS_CHECK等待海关审核 REFUND退款 DELETE删除 DRAFT 草稿)') index INT(11)"`
	Status         int       `json:"status" xorm:"not null default 0 comment('状态0未审核99已审核') index TINYINT(1)"`
	IsDel          int       `json:"is_del" xorm:"not null default 0 comment('是否删除1是0否') index TINYINT(1)"`
	TypeId         int       `json:"type_id" xorm:"not null default 0 comment('类别,1普通订单;') index INT(10)"`
	TypeIdAdmin    int       `json:"type_id_admin" xorm:"not null default 0 comment('类别,1普通订单;后台设置') index INT(10)"`
	TypeIdSub      int       `json:"type_id_sub" xorm:"not null default 0 comment('其他类别') INT(11)"`
	VatFee         int64     `json:"vat_fee" xorm:"not null default 0 comment('增值税费') BIGINT(20)"`
	SalesFee       int64     `json:"sales_fee" xorm:"not null default 0 comment('消费税') BIGINT(20)"`
	AmountFreight  int64     `json:"amount_freight" xorm:"not null default 0 comment('物流费用') BIGINT(20)"`
	AmountDiscount int64     `json:"amount_discount" xorm:"not null default 0 comment('折扣金额') BIGINT(20)"`
	AmountGoods    int64     `json:"amount_goods" xorm:"not null default 0 comment('商品总金额') BIGINT(20)"`
	AmountOther    int64     `json:"amount_other" xorm:"not null comment('其他价格费用') BIGINT(20)"`
	AmountTax      int64     `json:"amount_tax" xorm:"not null default 0 comment('税费') BIGINT(20)"`
	AmountOrder    int64     `json:"amount_order" xorm:"not null default 0 comment('订单总额') BIGINT(20)"`
	AmountPayment  int64     `json:"amount_payment" xorm:"not null default 0 comment('支付总额,已付款金额(实际付款金额)') BIGINT(20)"`
	Total          int       `json:"total" xorm:"not null default 0 comment('总数量') INT(10)"`
	TotalNoReceipt int       `json:"total_no_receipt" xorm:"not null default 0 comment('未收货数量') INT(10)"`
	Sid            int       `json:"sid" xorm:"not null default 0 comment('供应商ID') index INT(11)"`
	WarehouseId    int       `json:"warehouse_id" xorm:"not null default 0 comment('仓库ID') INT(10)"`
	StoreId        int       `json:"store_id" xorm:"not null default 0 comment('店铺ID') INT(10)"`
	ExpressNo      string    `json:"express_no" xorm:"not null default '' comment('物流单号,运送单号') CHAR(50)"`
	ExpressId      int       `json:"express_id" xorm:"not null default 0 comment('快递公司id') INT(10)"`
	Remark         string    `json:"remark" xorm:"comment('备注用户自己看') VARCHAR(255)"`
	RemarkAdmin    string    `json:"remark_admin" xorm:"comment('备注客服自己看') VARCHAR(255)"`
	GmtCreate      time.Time `json:"gmt_create" xorm:"default 'CURRENT_TIMESTAMP' comment('下单时间') index TIMESTAMP"`
	GmtModified    time.Time `json:"gmt_modified" xorm:"default 'CURRENT_TIMESTAMP' comment('最后更新时间') TIMESTAMP"`
	TimeSuccess    time.Time `json:"time_success" xorm:"comment('订单完成时间(整个订单完成，交易完成)') TIMESTAMP"`
	TimeCheckAdmin time.Time `json:"time_check_admin" xorm:"comment('客服审核时间') TIMESTAMP"`
	TimeCheck      time.Time `json:"time_check" xorm:"comment('审核时间，海关审核时间') TIMESTAMP"`
	TimeReceipt    time.Time `json:"time_receipt" xorm:"comment('收货时间') TIMESTAMP"`
	Declare        int       `json:"declare" xorm:"not null default 0 comment('DEFAULT未申报 NOT_ALLOW禁止申报 PORT_ACCEPT申报中 SUCCESS申报成功 FAIL申报失败 WARING申报异常;总订单时SUCCESS表示本订单已全部添加完成') INT(11)"`
	DeclareMsg     string    `json:"declare_msg" xorm:"comment('申报信息') VARCHAR(200)"`
	DeclareTime    time.Time `json:"declare_time" xorm:"comment('申报时间') TIMESTAMP"`
	IsSendTime     time.Time `json:"is_send_time" xorm:"comment('发货动作时间') TIMESTAMP"`
	IsSend         int       `json:"is_send" xorm:"not null default 0 comment('是否发货1是0否') index TINYINT(1)"`
	IsRefund       int       `json:"is_refund" xorm:"not null default 0 comment('是否退款') index TINYINT(1)"`
	IsReturn       int       `json:"is_return" xorm:"not null default 0 comment('退货1是0否') TINYINT(1)"`
	IsExchange     int       `json:"is_exchange" xorm:"not null default 0 comment('换货1是0否') TINYINT(1)"`
	OrderIdFrom    int       `json:"order_id_from" xorm:"not null default 0 comment('來自哪个ID，修改价格前ID') INT(10)"`
	OrderIdFromApi int       `json:"order_id_from_api" xorm:"not null default 0 comment('接口订单id') INT(10)"`
	OrderIdMaster  int       `json:"order_id_master" xorm:"not null default 0 comment('总订单号ID') INT(10)"`
	OrderNoMaster  string    `json:"order_no_master" xorm:"comment('总订单号') index CHAR(32)"`
	SidFrom        int       `json:"sid_from" xorm:"not null default 0 comment('货源商家id') INT(10)"`
	PayTime        time.Time `json:"pay_time" xorm:"comment('客户购买时间') TIMESTAMP"`
	PayId          int       `json:"pay_id" xorm:"not null default 0 comment('支付ID') INT(11)"`
	PayNo          string    `json:"pay_no" xorm:"comment('支付单号') CHAR(50)"`
	IsPaid         int       `json:"is_paid" xorm:"not null default 0 comment('是否已支付') TINYINT(1)"`
	IsPaidSystem   int       `json:"is_paid_system" xorm:"not null default 0 comment('是否已支付(系统自动)') TINYINT(1)"`
	TimePaidSystem time.Time `json:"time_paid_system" xorm:"comment('系统支付时间') DATETIME"`
	ExchangeRate   int64     `json:"exchange_rate" xorm:"not null default 0 comment('汇率') BIGINT(20)"`
	CurrencyMark   string    `json:"currency_mark" xorm:"comment('币制') CHAR(3)"`
	GetId          int       `json:"get_id" xorm:"not null default 0 comment('优惠券') INT(11)"`
	UseWallet      int64     `json:"use_wallet" xorm:"not null default 0 comment('使用钱包') BIGINT(20)"`
	UseCredit      int64     `json:"use_credit" xorm:"not null default 0 comment('使用积分') BIGINT(20)"`
}

//初始化
func NewOrder() *Order {
	return new(Order)
}
