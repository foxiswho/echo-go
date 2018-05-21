package order_service

import (
	"github.com/foxiswho/echo-go/dao/order_dao"
	"github.com/foxiswho/echo-go/models"
	"github.com/foxiswho/echo-go/util"
	"github.com/foxiswho/echo-go/consts/order_consts"
	"time"
	"github.com/foxiswho/echo-go/module/sn"
	"fmt"
)

//创建订单前格式化
type CreateOrderFormat struct {
	OrderGoodsData []*models.OrderGoodsData //数据集合
	create         *models.OrderCollectDate
	User           *models.User
	Options        *options
}

func NewCreateOrderFormat() *CreateOrderFormat {
	return new(CreateOrderFormat)
}

func (s *CreateOrderFormat) SetOrderGoodsData(data []*models.OrderGoodsData) {
	s.OrderGoodsData = data
}
func (s *CreateOrderFormat) SetUser(user *models.User) {
	s.User = user
}

func (s *CreateOrderFormat) SetOptions(options *options) {
	s.Options = options
}

func (s *CreateOrderFormat) setOrderGoods(data []*models.OrderGoods) {
	s.create.OrderGoods = data
}

func (s *CreateOrderFormat) Process() (*models.Order, error) {
	if s.OrderGoodsData == nil {
		return nil, util.NewError("商品数据不能为空")
	}
	if len(s.OrderGoodsData) == 0 {
		return nil, util.NewError("商品数据不能为空")
	}
	s.Options = newOptions()
	s.create = models.NewOrderCollectDate()
	s.create.Order = models.NewOrder()
	s.create.OrderExt = models.NewOrderExt()
	s.create.OrderConsignee = models.NewOrderConsignee()
	s.processGoods()
	s.processOrderData()
	s.processOptions()
	//
	order_create := order_dao.NewCreateOrder()
	order_create.SetOrderCollectDate(s.create)

	return order_create.Process()
}

//处理商品数据
func (s *CreateOrderFormat) processGoods() {
	count := len(s.OrderGoodsData)
	fmt.Println("OrderGoodsData=>", count, s.OrderGoodsData)
	//商品数据
	OrderGoods := make([]*models.OrderGoods, count)
	OrderGoodsStructure := make([]*models.OrderGoodsStructure, count)
	for key, goods := range s.OrderGoodsData {
		fmt.Println("key=>goods", key, goods.Goods)
		fmt.Println("key=>GoodsPrice", key, goods.GoodsPrice)
		if key == 0 {
			//对订单的仓库，和供应商赋值
			s.create.Order.Sid = goods.Goods.Sid
			s.create.Order.WarehouseId = goods.Goods.WarehouseId
		}
		order_goods := models.NewOrderGoods()
		order_goods.Num = goods.Num
		order_goods.Number = goods.Number
		order_goods.NumUnit = goods.NumUnit
		order_goods.NumTotal = order_goods.NumUnit * order_goods.Num
		order_goods.Title = goods.Title
		order_goods.Sid = goods.Sid
		order_goods.ProductId = goods.ProductId
		order_goods.WarehouseId = goods.WarehouseId
		order_goods.GoodsId = goods.Goods.Id
		order_goods.MarkId = goods.MarkId
		order_goods.Price = goods.Price
		order_goods.PriceShop = goods.PriceShop
		order_goods.Amount = int64(order_goods.Num) * order_goods.Price
		OrderGoods[key] = order_goods
		//组合商品数据 略,不处理组合商品
		//
		goods_structure := models.NewOrderGoodsStructure()
		goods_structure.Num = order_goods.Num
		goods_structure.Number = order_goods.Number
		goods_structure.NumUnit = order_goods.NumUnit
		goods_structure.NumTotal = order_goods.NumTotal
		goods_structure.Title = order_goods.Title
		goods_structure.Sid = order_goods.Sid
		goods_structure.ProductId = order_goods.ProductId
		goods_structure.WarehouseId = order_goods.WarehouseId
		goods_structure.GoodsId = order_goods.GoodsId
		goods_structure.Price = order_goods.Price
		goods_structure.PriceShop = order_goods.PriceShop
		goods_structure.Amount = order_goods.Amount
		goods_structure.ParentId = 0
		OrderGoodsStructure[key] = goods_structure
	}
	//fmt.Println("s.CreateOrder.OrderGoods=>", s.Order)
	fmt.Println("OrderGoods=>", len(OrderGoods), OrderGoods)
	//for key2, ggg := range OrderGoods {
	//	fmt.Println("key2,ggg=====")
	//	fmt.Println("key2,ggg=====")
	//	fmt.Println("key2,ggg=====")
	//	fmt.Println("key2,ggg=====", key2, ggg)
	//}
	//s.OrderGoods = make([]*models.OrderGoods, count)
	s.create.OrderGoods = OrderGoods
	s.create.OrderGoodsStructure = OrderGoodsStructure
	//for key2, ggg := range s.create.OrderGoods {
	//	fmt.Println("sss=====")
	//	fmt.Println("sss=====")
	//	fmt.Println("sss=====")
	//	fmt.Println("sss=====")
	//	fmt.Println("key2,ggg=====", key2, ggg)
	//}
}

//填充订单数据
func (s *CreateOrderFormat) processOrderData() {
	if s.create.Order.OrderNo == "" {
		s.create.Order.OrderNo = sn.MakeOrderNo()
	}
	if s.create.Order.PayId == 0 {
		s.create.Order.PayId = order_consts.PAY_ID_DEFAULT
	}
	s.create.Order.GmtCreate = time.Now()
	s.create.Order.GmtModified = s.create.Order.GmtCreate
	if s.create.Order.OrderStatus == 0 {
		s.create.Order.OrderStatus = int(order_consts.Order_Status_Not_Paid)
	}
	if s.create.Order.TypeId == 0 {
		s.create.Order.TypeId = int(order_consts.Order_Type_Id_Normal)
		s.create.Order.TypeIdAdmin = s.create.Order.TypeId
	}
	if s.create.Order.TypeIdAdmin == 0 {
		s.create.Order.TypeIdAdmin = s.create.Order.TypeId
	}
	s.create.Order.Status = 100
	s.create.Order.Uid = s.User.Id
}

//填充其他数据
func (s *CreateOrderFormat) processOptions() {
	if s.Options.OrderSn != "" {
		s.create.Order.OrderSn = s.Options.OrderSn
	}
}
