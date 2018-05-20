package goods_service

import (
	"github.com/foxiswho/echo-go/consts/goods_consts"
	"fmt"
)

func CreateGoodsTest() {
	create := NewCreate()
	create.Goods.Title = "德国Aptamil爱他美婴幼儿配方奶粉1+段(适合1岁以上宝宝)600g"
	create.Goods.Model = "1+段 600g"
	create.Goods.Number = "A000001"
	create.Goods.WarehouseId = 1
	create.Goods.Sid = 1
	create.Goods.ProductId = 1
	create.Goods.Status = 99
	create.Goods.IsOpen = 1
	create.Goods.TypeId = goods_consts.Type_Id_Normal
	create.Goods.CatId = 1
	create.Goods.BrandId = 1
	create.Goods.NumUnit = 1
	create.Goods.MarkId = int(goods_consts.MARK_ID_NORMAL)
	create.GoodsPrice.PriceShop = 1200000
	create.GoodsPrice.PriceMarket = 1400000
	create.GoodsPrice.NumMax = 999
	create.GoodsPrice.NumLeast = 1
	create.GoodsContent.Content = "内容"
	id, err := create.Process()
	fmt.Println("create goods id:", id)
	fmt.Println("create goods err:", err)
}
