package service

import (
	"context"
	"fmt"
	"go-store/pkg/mytools"
	"math/rand"

	"go-store/dao"
	"go-store/model"
	"go-store/pkg/e"
	"go-store/serializer"
	"net/http"
	"strconv"
	"time"
)

type OrderService struct {
	UserID     int     `json:"user_id" form:"user_id"`         // 用户 ID，不能为空
	ProductID  int     `json:"product_id" form:"product_id"`   // 商品 ID，不能为空
	MerchantID int     `json:"merchant_id" form:"merchant_id"` // 商家 ID，不能为空
	AddressID  int     `json:"address_id" form:"address_id"`   // 地址 ID，不能为空
	Num        int     `json:"num" form:"num"`                 // 商品数量
	OrderNum   int     `json:"order_num" form:"order_num"`     // 订单编号
	Type       uint    `json:"type" form:"type"`               // 订单类型，0 未支付，1 为已支付
	Money      float64 `json:"money" form:"money"`             // 订单金额
	Key        string  `json:"key" form:"key"`                 // 支付秘钥
	model.BasePage
}

func (s *OrderService) CreateOrderService(context context.Context, uid int) interface{} {
	var order *model.Order
	code := http.StatusOK
	orderDao := dao.NewOrderDao(context)
	order = &model.Order{
		UserID:     uid,
		ProductID:  s.ProductID,
		MerchantID: s.MerchantID,
		AddressID:  s.AddressID,
		Num:        s.Num,
		OrderNum:   s.OrderNum,
		Type:       0, // 默认未支付
		Money:      s.Money,
	}

	addressDao := dao.NewAddressDao(context)
	// 检查地址是否存在
	addr, err := addressDao.GetAddressByID(s.AddressID)
	if err != nil {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code) + "地址不存在",
			Error:  err.Error(),
		}
	}
	order.AddressID = int(addr.ID)

	// 检查商品是否存在
	productDao := dao.NewProductDao(context)
	_, err = productDao.GetProductByID(s.ProductID)
	if err != nil {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code) + "商品已下架",
			Error:  err.Error(),
		}
	}

	number := time.Now().Unix() + int64(rand.Intn(1000)) // 以当前时间戳作为订单编号+随机数
	order.OrderNum = int(number)
	err = orderDao.CreateOrder(order)
	if err != nil {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMSG(code),
	}
}

func (s *OrderService) GetOrderByID(ctx context.Context, uid int, oid string) serializer.Response {
	orderId, _ := strconv.Atoi(oid)
	code := http.StatusOK
	OrderDao := dao.NewOrderDao(ctx)
	order, err := OrderDao.GetOrderByID(orderId, uid)
	if err != nil {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    "查询订单不存在",
			Error:  err.Error(),
		}
	}
	if uid != order.UserID {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    "账户信息异常",
			Error:  err.Error(),
		}
	}
	// 订单地址
	AddressDao := dao.NewAddressDao(ctx)
	address, err := AddressDao.GetAddressByID(order.AddressID)
	if err != nil {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    "订单地址不存在",
			Error:  err.Error(),
		}
	}

	// 订单商品
	ProductDao := dao.NewProductDao(ctx)
	product, err := ProductDao.GetProductByID(order.ProductID)
	if err != nil {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    "订单商品不存在",
			Error:  err.Error(),
		}
	}

	// 用户信息
	UserDao := dao.NewUserDao(ctx)
	user, err := UserDao.GetUserByID(order.UserID)
	boss, err := UserDao.GetUserByID(order.MerchantID)
	if err != nil {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMSG(code),
		Data:   serializer.BuildOrder(order, product, address, user, boss),
	}
}

func (s *OrderService) ShowOrderService(ctx context.Context, uid int) serializer.Response {
	code := http.StatusOK
	if s.PageSize == 0 {
		s.PageSize = 15
	}

	orderDao := dao.NewOrderDao(ctx)
	condition := make(map[string]interface{})
	if s.Type == 2 { // 查询全部订单
		condition["type"] = s.Type
	} else if s.Type == 1 {
		condition["type"] = s.Type // 查询已付订单
	} else if s.Type == 0 {
		condition["type"] = s.Type // 查询未支付订单
	}
	condition["user_id"] = uid
	orderList, err, total := orderDao.ListOrderCondition(condition, s.BasePage)
	if err != nil {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListCarouse(serializer.BuildOrders(ctx, orderList), total)
}

func (s *OrderService) DeleteOrder(ctx context.Context, uid int, aid string) interface{} {
	orderDao := dao.NewOrderDao(ctx)
	addrID, _ := strconv.Atoi(aid)
	err := orderDao.DeleteOrderByID(addrID, uid)
	if err != nil {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    e.GetMSG(http.StatusBadRequest),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: http.StatusOK,
		Msg:    e.GetMSG(http.StatusOK),
	}
}

func (s *OrderService) PayDown(ctx context.Context, uid int, oid string) serializer.Response {
	orid, _ := strconv.Atoi(oid)
	mytools.Encrypt.SetKey(s.Key)
	code := http.StatusOK
	orderDao := dao.NewOrderDao(ctx)
	tx := orderDao.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			mytools.Logger.Infoln("panic occurred: ", r)
			code = http.StatusInternalServerError
		}
	}()

	// 封装错误处理函数
	handleError := func(err error, msg string) serializer.Response {
		tx.Rollback()
		mytools.Logger.Infoln("err: ", err)
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    msg,
			Error:  err.Error(),
		}
	}

	// 获取订单信息
	order, err := orderDao.GetOrderByID(orid, uid)
	if err != nil {
		return handleError(err, e.GetMSG(code))
	}

	// 计算总金额
	totalMoney := order.Money * float64(order.Num)

	userDao := dao.NewUserDao(ctx)
	userDaoDB := dao.NewUserDaoByDB(userDao.DB)

	// 处理用户扣款
	user, err := userDao.GetUserByID(order.UserID)
	if err != nil {
		return handleError(err, e.GetMSG(code))
	}

	userMoney, err := decryptMoney(user.Money)
	if err != nil {
		return handleError(err, "用户金额解密异常")
	}

	if userMoney < totalMoney {
		tx.Rollback()
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    e.GetMSG(http.StatusBadRequest),
			Error:  "余额不足",
		}
	}

	newUserMoney := userMoney - totalMoney
	encryptedMoney, err := encryptMoney(newUserMoney)
	if err != nil {
		return handleError(err, "用户金额加密异常")
	}
	user.Money = encryptedMoney

	if err := userDaoDB.UpdateUserByID(uid, user); err != nil {
		return handleError(err, e.GetMSG(code))
	}

	// 处理商家加钱
	boss, err := userDao.GetUserByID(order.MerchantID)
	if err != nil {
		return handleError(err, e.GetMSG(code))
	}

	bossMoney, err := decryptMoney(boss.Money)
	if err != nil {
		return handleError(err, "商家金额解密异常")
	}

	newBossMoney := bossMoney + totalMoney
	encryptedBossMoney, err := encryptMoney(newBossMoney)
	if err != nil {
		return handleError(err, "商家金额加密异常")
	}
	boss.Money = encryptedBossMoney

	if err := userDaoDB.UpdateUserByID(order.MerchantID, boss); err != nil {
		return handleError(err, e.GetMSG(code))
	}

	// 处理商品库存
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductByID(order.ProductID)
	if err != nil {
		return handleError(err, e.GetMSG(code))
	}

	product.Number -= order.Num
	fmt.Println("order.Num", order.Num, "product.Number", product.Number)
	if product.Number < 0 {
		tx.Rollback()
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    e.GetMSG(http.StatusBadRequest),
			Error:  "库存不足",
		}
	}

	if err := productDao.UpdateProduct(order.ProductID, product); err != nil {
		return handleError(err, "库存更新失败")
	}

	// 更新订单状态
	if order.Type == 1 {
		tx.Rollback()
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    e.GetMSG(http.StatusBadRequest),
			Error:  "订单已完成交易,无需再次交易",
		}
	}
	order.Type = 1
	fmt.Println(order.Type)
	if err := orderDao.UpdateOrderByID(orid, uid, order); err != nil {
		tx.Rollback()
		return handleError(err, "订单更新失败")
	}
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return handleError(err, "订单交易失败")
	}

	return serializer.Response{
		Status: code,
		Msg:    "订单完成交易",
	}
}

// decryptMoney 解密金额
func decryptMoney(encryptedMoney string) (float64, error) {
	moneyStr, err := mytools.Encrypt.MyDecrypt(encryptedMoney)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(moneyStr, 64)
}

// encryptMoney 加密金额
func encryptMoney(money float64) (string, error) {
	finMoney := fmt.Sprintf("%.2f", money)
	return mytools.Encrypt.MyEncrypt(finMoney)
}
