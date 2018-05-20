package goods_dao

import (
	"github.com/foxiswho/echo-go/models"
	"github.com/foxiswho/echo-go/module/db"
	"github.com/foxiswho/echo-go/util"
	"github.com/foxiswho/echo-go/module/goods"
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
	engine := db.DB().Engine
	id, err := engine.InsertOne(s.Goods)
	if err != nil {
		return 0, util.NewError("商品 保存错误")
	}
	s.Goods.Id = int(id)
	//更新标志
	engine.Where("mark=?", goods.GetMarkByGoods(s.Goods)).Update(models.NewGoods())
	//
	s.GoodsPrice.Id = s.Goods.Id
	_, err = engine.InsertOne(s.GoodsPrice)
	if err != nil {
		return 0, util.NewError("商品价格 保存错误")
	}
	s.GoodsContent.Id = s.Goods.Id
	_, err = engine.InsertOne(s.GoodsContent)
	if err != nil {
		return 0, util.NewError("商品内容 保存错误")
	}

	return s.Goods.Id, nil
}
