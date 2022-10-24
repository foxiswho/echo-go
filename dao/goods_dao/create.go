package goods_dao

import (
	"fmt"

	"github.com/foxiswho/echo-go/models"
	"github.com/foxiswho/echo-go/module/db"
	"github.com/foxiswho/echo-go/module/goods"
	"github.com/foxiswho/echo-go/module/log"
	"github.com/foxiswho/echo-go/util"
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
	affected, err := engine.InsertOne(s.Goods)
	if err != nil {
		log.Debugf("s.Goods", err.Error())
		return 0, util.NewError("商品 保存错误")
	}
	log.Debugf("s.Goods affected ", affected)
	//更新标志
	mark := models.NewGoods()
	mark.Mark = goods.GetMarkByGoods(s.Goods)
	fmt.Println("update mark ", mark.Mark)
	up, err := engine.ID(s.Goods.Id).Update(mark)
	fmt.Println("UPDATE", up)
	fmt.Println("UPDATE err", err)
	//
	s.GoodsPrice.Id = s.Goods.Id
	_, err = engine.InsertOne(s.GoodsPrice)
	if err != nil {
		log.Debugf("s.GoodsPrice", err.Error())
		return 0, util.NewError("商品价格 保存错误")
	}
	s.GoodsContent.Id = s.Goods.Id
	_, err = engine.InsertOne(s.GoodsContent)
	if err != nil {
		log.Debugf("s.GoodsContent", err.Error())
		return 0, util.NewError("商品内容 保存错误")
	}

	return s.Goods.Id, nil
}
