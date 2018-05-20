package order_consts

//支付宝
const (
	PAY_ID_DEFAULT = 100001
	//
)

//订单状态
type OrderStatus int

const (
	Order_Status_Not_Paid           OrderStatus = 101010
	Order_Status_Wait_Check         OrderStatus = 101020
	Order_Status_No_Check           OrderStatus = 101030
	Order_Status_Wait_Customs_Check OrderStatus = 101040
	Order_Status_Wait_Send          OrderStatus = 101050
	Order_Status_Send               OrderStatus = 101060
	Order_Status_Receipt            OrderStatus = 101070
	Order_Status_Success            OrderStatus = 101100
	Order_Status_Cancel             OrderStatus = 101110
	Order_Status_Delete             OrderStatus = 101120
)

type OrderType int

const (
	Order_Type_Id_Normal       OrderType = 102001 //普通订单
	Order_Type_Id_Direct_Mail  OrderType = 102002 //直邮
	Order_Type_Id_stock        OrderType = 102003 //现货
	Order_Type_Id_quick_order  OrderType = 102010 //快速下单
	Order_Type_Id_import_order OrderType = 102011 //导入下单
	Order_Type_Id_2b           OrderType = 102100 //批量b2b订单
)
