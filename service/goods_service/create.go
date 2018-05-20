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
	fmt.Println("s.Goods", s.Goods)
	fmt.Println("s.GoodsPrice", s.GoodsPrice)
	fmt.Println("s.GoodsContent", s.GoodsContent)
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
	return id, nil
}

func CreateGoodsxxxxxx() {
	create := NewCreate()
	goods := models.NewGoods()
	goods.Title = "德国Aptamil爱他美婴幼儿配方奶粉1+段(适合1岁以上宝宝)600g"
	goods.Model = "1+段 600g"
	goods.Number = "A000001"
	goods.WarehouseId = 1
	goods.Sid = 1
	goods.ProductId = 1
	goods.Status = 99
	goods.IsOpen = 1
	goods.TypeId = goods_consts.Type_Id_Normal
	goods.CatId = 1
	goods.BrandId = 1
	goods.NumUnit = 1
	goods.MarkId = int(goods_consts.MARK_ID_NORMAL)
	create.Goods = goods
	create.GoodsPrice = models.NewGoodsPrice()
	create.GoodsPrice.PriceShop = 1200000
	create.GoodsPrice.PriceMarket = 1400000
	create.GoodsPrice.NumMax = 999
	create.GoodsPrice.NumLeast = 1
	create.GoodsContent = models.NewGoodsContent()
	create.GoodsContent.Content = "内容"
	fmt.Println("create.Goods", create.Goods)
	fmt.Println("create.GoodsPrice", create.GoodsPrice)
	fmt.Println("create.GoodsContent", create.GoodsContent)
	id, err := create.Process()
	fmt.Println("create goods id:", id)
	fmt.Println("create goods err:", err)
}
