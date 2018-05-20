package goods_consts

type MarkId int

//标志类型
const (
	//正常商品
	MARK_ID_NORMAL MarkId = 201001 + iota
	//组合商品
	MARK_ID_COMBINED
)
