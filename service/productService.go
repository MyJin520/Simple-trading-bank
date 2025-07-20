package service

import (
	"context"
	"fmt"
	"go-store/dao"
	"go-store/model"
	"go-store/pkg/e"
	"go-store/pkg/mytools"
	"go-store/serializer"
	"mime/multipart"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type ProductService struct {
	ID             int    `form:"id" json:"id"`                         // 商品ID
	Title          string `form:"title" json:"title"`                   // 商品标题
	Name           string `form:"name" json:"name"`                     // 商品名称
	Info           string `form:"info" json:"info"`                     // 商品简介
	ImagePath      string `form:"img_path" json:"img_path"`             // 以第一张图为展示图
	CategoryID     int    `form:"category_id" json:"category_id"`       // 类别ID
	Price          string `form:"price" json:"price"`                   // 价格
	DiscountPrice  string `form:"discount_price" json:"discount_price"` // 折扣价格
	OnSale         bool   `form:"on_sale" json:"on_sale"`               // 是否上架
	Num            int    `form:"num" json:"num"`                       // 库存
	model.BasePage        // 分页
}

// CreateProductService 创建商品服务
func (ps *ProductService) CreateProductService(ctx context.Context, uid int, files []*multipart.FileHeader) serializer.Response {
	var (
		boss *model.User
		err  error
	)
	code := http.StatusOK

	userDao := dao.NewUserDao(ctx)
	boss, _ = userDao.GetUserByID(uid)

	//以第一张做为封面图片
	tmp, _ := files[0].Open()
	path, err := UploadProductLocalStatic(tmp, uid, ps.Name)
	if err != nil {
		mytools.Logger.Infoln("err: ", err)
		code = 50001
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}

	// 创建一个新的商品模型实例，用于存储商品信息到数据库
	product := &model.Product{
		// 商品名称，从 ProductService 结构体中获取
		Name: ps.Name,
		// 商品类别，从 ProductService 结构体中获取
		CategoryID: ps.CategoryID,
		// 商品标题，从 ProductService 结构体中获取
		Title: ps.Title,
		// 商品简介，从 ProductService 结构体中获取
		Info: ps.Info,
		// 商品图片路径，使用上传第一张图片得到的路径
		ImgPath: path,
		// 商品价格，从 ProductService 结构体中获取
		Price: ps.Price,
		// 商品折扣价格，从 ProductService 结构体中获取，注意此处原代码字段名可能拼写错误，应为 DiscountPrice
		DiscountPerice: ps.DiscountPrice,
		// 商品库存数量，从 ProductService 结构体中获取
		Number: ps.Num,
		// 商家 ID，由传入的用户 ID 确定
		MerchantID: uid,
		// 商家名称，从查询到的用户信息中获取
		MerchantName: boss.UserName,
		// 商家头像路径，从查询到的用户信息中获取
		MerchantAvatar: boss.Avatar,
		// 商品是否上架，默认为上架状态
		OnSale: true,
	}

	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		code = http.StatusBadRequest
		mytools.Logger.Infoln("err: ", err)
		return serializer.Response{
			Status: code,
			Data:   e.GetMSG(code),
			Error:  err.Error(),
		}
	}
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		fmt.Println("filename:", file.Filename, len(files))
		num := strconv.Itoa(index)
		productImagDao := dao.NewProductImageDaoByID(productDao.DB)
		tmp, _ := file.Open()
		path, _ = UploadProductLocalStatic(tmp, uid, ps.Name+num+time.Now().Format("2006-01-02 15:04:05"))

		productImg := &model.ProductImg{
			ProductID: int(product.ID),
			ImgPath:   path,
		}
		err = productImagDao.CreateProductImage(productImg)
		if err != nil {
			mytools.Logger.Infoln("err: ", err)
			code = 50002
			return serializer.Response{
				Status: code,
				Msg:    e.GetMSG(code),
				Error:  err.Error(),
			}
		}
		wg.Done()
	}
	wg.Wait()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMSG(code),
		Data:   serializer.BuildProduct(product),
	}
}

func (ps *ProductService) ListProduct(ctx context.Context) serializer.Response {
	var (
		products []*model.Product
		err      error
	)
	code := http.StatusOK
	if ps.PageSize == 0 {
		ps.PageSize = 15
	}
	condition := make(map[string]interface{})
	if ps.CategoryID != 0 {
		condition["category_id"] = ps.CategoryID
	}
	productDao := dao.NewProductDao(ctx)
	total, err := productDao.CountProductByCondition(condition)
	if err != nil {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.ListProductByCondition(condition, ps.BasePage)
		wg.Done()
	}()
	wg.Wait()
	return serializer.BuildListCarouse(serializer.BuildProducts(products), total)
}

func (ps *ProductService) FindProduct(ctx context.Context) serializer.Response {
	code := http.StatusOK
	if ps.PageSize == 0 {
		ps.PageSize = 15
	}
	productDao := dao.NewProductDao(ctx)
	product, count, err := productDao.FindProduct(ps.Info, ps.BasePage)
	if err != nil {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListCarouse(serializer.BuildProducts(product), count)
}

func (ps *ProductService) ShowProduct(ctx context.Context, id string) serializer.Response {
	code := http.StatusOK
	pID, _ := strconv.Atoi(id)
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductByID(pID)
	if err != nil {
		code = http.StatusBadRequest
		mytools.Logger.Infoln("err: ", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMSG(code),
		Data:   serializer.BuildProduct(product),
	}
}
