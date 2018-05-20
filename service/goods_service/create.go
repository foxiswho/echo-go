package goods_service

import (
	"github.com/foxiswho/echo-go/models"
	"github.com/foxiswho/echo-go/util"
	"github.com/foxiswho/echo-go/dao/goods_dao"
	"github.com/foxiswho/echo-go/consts/goods_consts"
	"fmt"
)

type Create struct {
	*models.Goods
	*models.GoodsPrice
	*models.GoodsContent
}

func NewCreate() *Create {
	return new(Create)
}

func (s *Create) Process() (int, error) {
	if s.Goods.Title == "" {
		return 0, util.NewError("商品 名称不能为空")
	}
	goods_dao := goods_dao.NewCreate()
	goods_dao.Goods = s.Goods
	goods_dao.GoodsPrice = s.GoodsPrice
	goods_dao.GoodsContent = s.GoodsContent
	id, err := goods_dao.Process()
	if err != nil {
		return 0, err
	}
	return id, util.NewError("商品 名称不能为空")
}


func CreateGoodsxxxxxx() {
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
