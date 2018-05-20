package goods_service

import (
	"github.com/foxiswho/echo-go/models"
	"github.com/foxiswho/echo-go/util"
	"github.com/foxiswho/echo-go/dao/goods_dao"
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
