
package service

import (
	"fmt"
	"github.com/foxiswho/shop-go/models"
	"github.com/foxiswho/shop-go/module/db"
	"github.com/foxiswho/shop-go/util"
)

type OrderExtService struct {

}

func NewOrderExtService() *OrderExtService {
	return new(OrderExtService)
}

//初始化列表
func order_extNewMakeDataArr() []models.OrderExt {
	return make([]models.OrderExt, 0)
}

//列表查询
func (s *OrderExtService) GetAll(where []*db.QueryCondition, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
	m := models.NewOrderExt()
	session := db.Filter(where)
	count, err := session.Count(m)
	if err != nil {
		fmt.Println(err)
		return nil, util.NewError(err.Error())
	}
	Query := db.Pagination(int(count), page, limit)
	if count == 0 {
		return Query, nil
	}

	session = db.Filter(where)
	if orderBy != "" {
		session.OrderBy(orderBy)
	}
	session.Limit(limit, Query.Offset)
	if len(fields) == 0 {
		session.AllCols()
	}
	data := order_extNewMakeDataArr()
	err = session.Find(&data)
	if err != nil {
		fmt.Println(err)
		return nil, util.NewError(err.Error())
	}
	Query.Data = make([]interface{}, len(data))
	for y, x := range data {
		Query.Data[y] = x
	}
	return Query, nil
}


// 获取 单条记录
func (s *OrderExtService) GetById(id int) (*models.OrderExt, error) {
    m:=new(models.OrderExt)
	m.Id = id
	ok, err := db.DB().Engine.Get(m)
    if err != nil {
        return nil, err
    }
    if !ok{
        return nil,util.NewError("数据不存在:"+err.Error())
    }
    return m, nil
}

// 删除 单条记录
func (s *OrderExtService) Delete(id int) (int64, error) {
	m:=new(models.OrderExt)
	m.Id = id
	num, err := db.DB().Engine.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}