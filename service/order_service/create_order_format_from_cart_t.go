package order_service

import (
	"github.com/foxiswho/echo-go/module/db"
	"github.com/foxiswho/echo-go/models"
	"github.com/foxiswho/echo-go/service"
	"fmt"
	"github.com/foxiswho/echo-go/consts/cart_consts"
)

func CreateOrderFormatFromCartT() {
	//
	uid := 1
	user, err := service.NewUserService().GetById(uid)
	if err != nil {
		fmt.Println("err", err)
	}
	//查询出指定商品数据
	goods_id := 1002
	goods, err := service.NewGoodsService().GetById(goods_id)
	if err != nil {
		fmt.Println("err", err)
	}
	goods_price, err := service.NewGoodsPriceService().GetById(goods_id)
	if err != nil {
		fmt.Println("err", err)
	}
	//购物车数据 组合
	cart := models.NewCart()
	engine := db.DB().Engine
	//查询是否存在
	count, err := engine.Where("uid=?", uid).Count(cart)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("count", count)
	if count == 0 {
		//保存到数据库
		cart.Num = 10
		cart.Price = goods_price.PriceShop
		cart.GoodsId = goods.Id
		cart.ProductId = goods.ProductId
		cart.WarehouseId = goods.WarehouseId
		cart.Sid = goods.Sid
		cart.TypeId = cart_consts.Type_Id_Normal
		cart.Amount = int64(cart.Num) * cart.Price
		cart.Uid = uid

		affected, err := engine.Insert(cart)
		if err != nil {
			fmt.Println("err", err)
		}
		fmt.Println("affected", affected)
	}
	///////////////////////////////////////////////////////
	//获取购物车数据
	carts := make([]models.Cart, 0)
	err = engine.Where("uid=?", uid).Find(&carts)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println("carts", carts)
	//填充数据
	cart_service := NewCreateOrderFormatFromCart()
	cart_service.SetCart(carts)
	order_goods, err := cart_service.Process()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("cart_service=>order_goods", order_goods)
	order_create := NewCreateOrderFormat()
	order_create.SetOrderGoodsData(order_goods)
	order_create.SetUser(user)
	order, err := order_create.Process()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("order", order)
}
