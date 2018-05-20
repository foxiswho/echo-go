package goods_consts

type MarkId int

//标志类型
const (
	//正常商品
	MARK_ID_NORMAL MarkId = 201001 + iota
	//组合商品
	MARK_ID_COMBINED
)

const (
	Type_Id_Normal         = 202001 //普通订单
	Type_Id_Direct_Mail    = 202002 //直邮
	Type_Id_goods_in_stock = 202003 //现货
	Type_Id_bonded_area    = 202004 //保税区
)
