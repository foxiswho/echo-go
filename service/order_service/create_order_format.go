package order_service

import (
	"github.com/foxiswho/echo-go/dao/order_dao"
	"github.com/foxiswho/echo-go/models"
	"github.com/foxiswho/echo-go/util"
	"github.com/foxiswho/echo-go/consts/order_consts"
	"time"
	"github.com/foxiswho/echo-go/module/sn"
)

//创建订单前格式化
type CreateOrderFormat struct {
	OrderGoodsData []*models.OrderGoodsData //数据集合
	createOrder    *order_dao.CreateOrder
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

func (s *CreateOrderFormat) Process() (*models.Order, error) {
	if s.OrderGoodsData == nil {
		return nil, util.NewError("商品数据不能为空")
	}
	if len(s.OrderGoodsData) == 0 {
		return nil, util.NewError("商品数据不能为空")
	}
	s.processGoods()
	s.processOrderData()
	s.processOptions()

	return s.createOrder.Process()
}

//处理商品数据
func (s *CreateOrderFormat) processGoods() {
	//商品数据
	s.createOrder.OrderGoods = make([]*models.OrderGoods, 0)
	s.createOrder.OrderGoodsStructure = make([]*models.OrderGoodsStructure, 0)
	for key, goods := range s.OrderGoodsData {
		if key == 0 {
			//对订单的仓库，和供应商赋值
			s.createOrder.Order.Sid = goods.Goods.Sid
			s.createOrder.Order.WarehouseId = goods.Goods.WarehouseId
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
		s.createOrder.OrderGoods = append(s.createOrder.OrderGoods, order_goods)
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
		s.createOrder.OrderGoodsStructure = append(s.createOrder.OrderGoodsStructure, goods_structure)
	}
}

//填充订单数据
func (s *CreateOrderFormat) processOrderData() {
	if s.createOrder.Order.OrderNo == "" {
		s.createOrder.Order.OrderNo = sn.MakeOrderNo()
	}
	if s.createOrder.Order.PayId == 0 {
		s.createOrder.Order.PayId = order_consts.PAY_ID_DEFAULT
	}
	s.createOrder.Order.GmtCreate = time.Now()
	s.createOrder.Order.GmtModified = s.createOrder.Order.GmtCreate
	if s.createOrder.Order.OrderStatus == 0 {
		s.createOrder.Order.OrderStatus = int(order_consts.Order_Status_Not_Paid)
	}
	if s.createOrder.Order.TypeId == 0 {
		s.createOrder.Order.TypeId = int(order_consts.Order_Type_Id_Normal)
		s.createOrder.Order.TypeIdAdmin = s.createOrder.Order.TypeId
	}
	if s.createOrder.Order.TypeIdAdmin == 0 {
		s.createOrder.Order.TypeIdAdmin = s.createOrder.Order.TypeId
	}
	s.createOrder.Order.Status = 100
	s.createOrder.Order.Uid = s.User.Id
}

//填充其他数据
func (s *CreateOrderFormat) processOptions() {
	if s.Options.OrderSn != "" {
		s.createOrder.Order.OrderSn = s.Options.OrderSn
	}
}
