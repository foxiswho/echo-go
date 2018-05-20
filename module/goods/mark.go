package goods

import (
	"github.com/foxiswho/echo-go/consts/goods_consts"
	"github.com/foxiswho/echo-go/models"
)

//根据标志类型，产品ID，仓库ID，供应商ID 获取标志
func GetMark(mark_id goods_consts.MarkId, product_id, warehouse_id, sid int) string {
	goods := new(models.Goods)
	goods.MarkId = int(mark_id)
	goods.ProductId = product_id
	goods.WarehouseId = warehouse_id
	goods.Sid = sid
	return GetMarkByGoods(goods)
}

//根据商品数据获取唯一标志
func GetMarkByGoods(goods *models.Goods) string {
	if int(goods_consts.MARK_ID_NORMAL) == goods.MarkId {
		return string(goods.MarkId) + "-" + string(goods.Id)
	} else {
		return string(goods.MarkId) + "-" + string(goods.ProductId) + "-" + string(goods.WarehouseId) + "-" + string(goods.Sid)
	}
}
