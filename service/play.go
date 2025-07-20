package service

type OrderPayService struct {
	OrderID   int     `form:"order_id" json:"order_id"`
	Money     float64 `json:"money" form:"money"`
	OrderNo   string  `json:"order_no" form:"order_no"`
	ProductID int     `json:"product_id" form:"product_id"`
	PayTime   string  `json:"pay_time" form:"pay_time"`
	Sing      string  `json:"sing" form:"sing"`
	BossID    int     `json:"boss_id" form:"boss_id"`
	UserID    int     `json:"user_id" form:"user_id"`
	Num       int     `json:"num" form:"num"`
	Key       string  `json:"key" form:"key"` // 支付秘钥
}

//func (s *OrderPayService) PayDown(ctx context.Context, uid int, oid string) serializer.Response {
//	orid,_ := strconv.Atoi(oid)
//	mytools.Encrypt.SetKey(s.Key)
//	code := http.StatusOK
//	orderDao := dao.NewOrderDao(ctx)
//	tx := orderDao.Begin()
//	defer func() {
//		if r := recover(); r != nil {
//			tx.Rollback()
//			mytools.Logger.Infoln("panic occurred: ", r)
//			code = http.StatusInternalServerError
//		}
//	}()
//
//	// 封装错误处理函数
//	handleError := func(err error, msg string) serializer.Response {
//		tx.Rollback()
//		mytools.Logger.Infoln("err: ", err)
//		code = http.StatusBadRequest
//		return serializer.Response{
//			Status: code,
//			Msg:    msg,
//			Error:  err.Error(),
//		}
//	}
//
//	// 获取订单信息
//	order, err := orderDao.GetOrderByID(orid, uid)
//	if err != nil {
//		return handleError(err, e.GetMSG(code))
//	}
//
//	// 计算总金额
//	totalMoney := order.Money * float64(order.Num)
//
//	userDao := dao.NewUserDao(ctx)
//	userDaoDB := dao.NewUserDaoByDB(userDao.DB)
//
//	// 处理用户扣款
//	user, err := userDao.GetUserByID(order.UserID)
//	if err != nil {
//		return handleError(err, e.GetMSG(code))
//	}
//
//	userMoney, err := decryptMoney(user.Money)
//	if err != nil {
//		return handleError(err, "金额解密异常")
//	}
//
//	if userMoney < totalMoney {
//		tx.Rollback()
//		return serializer.Response{
//			Status: http.StatusBadRequest,
//			Msg:    e.GetMSG(http.StatusBadRequest),
//			Error:  "余额不足",
//		}
//	}
//
//	newUserMoney := userMoney - totalMoney
//	encryptedMoney, err := encryptMoney(newUserMoney)
//	if err != nil {
//		return handleError(err, "金额加密异常")
//	}
//	user.Money = encryptedMoney
//
//	if err := userDaoDB.UpdateUserByID(uid, user); err != nil {
//		return handleError(err, e.GetMSG(code))
//	}
//
//	// 处理商家加钱
//	boss, err := userDao.GetUserByID(order.MerchantID)
//	if err != nil {
//		return handleError(err, e.GetMSG(code))
//	}
//
//	bossMoney, err := decryptMoney(boss.Money)
//	if err != nil {
//		return handleError(err, "金额解密异常")
//	}
//
//	newBossMoney := bossMoney + totalMoney
//	encryptedBossMoney, err := encryptMoney(newBossMoney)
//	if err != nil {
//		return handleError(err, "金额加密异常")
//	}
//	boss.Money = encryptedBossMoney
//
//	if err := userDaoDB.UpdateUserByID(order.MerchantID, boss); err != nil {
//		return handleError(err, e.GetMSG(code))
//	}
//
//	// 处理商品库存
//	productDao := dao.NewProductDao(ctx)
//	product, err := productDao.GetProductByID(order.ProductID)
//	if err != nil {
//		return handleError(err, e.GetMSG(code))
//	}
//
//	product.Number -= order.Num
//	if product.Number < 0 {
//		tx.Rollback()
//		return serializer.Response{
//			Status: http.StatusBadRequest,
//			Msg:    e.GetMSG(http.StatusBadRequest),
//			Error:  "库存不足",
//		}
//	}
//
//	if err := productDao.UpdateProduct(order.ProductID, product); err != nil {
//		return handleError(err, "库存更新失败")
//	}
//
//	// 更新订单状态
//	order.Type = 1
//
//	// 提交事务
//	if err := tx.Commit().Error; err != nil {
//		return handleError(err, "订单交易失败")
//	}
//
//	return serializer.Response{
//		Status: code,
//		Msg:    "订单完成交易",
//	}
//}
//
//// decryptMoney 解密金额
//func decryptMoney(encryptedMoney string) (float64, error) {
//	moneyStr, err := mytools.Encrypt.MyDecrypt(encryptedMoney)
//	if err != nil {
//		return 0, err
//	}
//	return strconv.ParseFloat(moneyStr, 64)
//}
//
//// encryptMoney 加密金额
//func encryptMoney(money float64) (string, error) {
//	finMoney := fmt.Sprintf("%.2f", money)
//	return mytools.Encrypt.MyEncrypt(finMoney)
//}
