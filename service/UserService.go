package service

import (
	"context"
	"fmt"
	"go-store/conf"
	"go-store/dao"
	"go-store/model"
	"go-store/pkg/e"
	"go-store/pkg/mytools"
	"go-store/serializer"
	"gopkg.in/mail.v2"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"strings"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"` // 添加 form 标签
	UserName string `json:"user_name" form:"user_name"` // 添加 form 标签
	PassWord string `json:"password" form:"password"`   // 添加 form 标签
	Email    string `json:"email" form:"email"`         // 添加 form 标签
	Key      string `json:"key" form:"key"`
	Money    string `json:"money" form:"money"`
}

type EmailService struct {
	Email         string `json:"email" form:"email"`
	PassWord      string `json:"password" form:"password"`
	OperationType int    `json:"operation_type" gorm:"operation_type"` // 1.绑定邮箱，2、解绑邮箱，3、修改密码
}

type VerifyEmailService struct {
	Email         string `json:"email" form:"email"`
	PassWord      string `json:"password" form:"password"`
	OperationType int    `json:"operation_type" gorm:"operation_type"` // 1.绑定邮箱，2、解绑邮箱，3、修改密码
}

type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
}

func (s *ShowMoneyService) ShowMoney(ctx context.Context, id int) serializer.Response {
	// 初始化响应状态码
	code := http.StatusOK
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserByID(id)
	if err != nil {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUserMoney(user, user.Money, s.Key),
		Msg:    e.GetMSG(code),
	}
}

// SendEmail 发送验证邮件
func (s *EmailService) SendEmail(ctx context.Context, id int) serializer.Response {
	// 初始化响应状态码
	code := http.StatusOK

	var (
		address string
		notice  *model.Notice // 用于存储通知模板信息
	)

	// 生成邮箱验证 Token
	token, err := mytools.CreateEmailToken(id, s.OperationType, s.Email, s.PassWord)
	if err != nil {
		// 生成 Token 失败
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}

	// 创建 NoticeDao 实例，用于查询通知模板
	noticeDao := dao.NewNoticeDao(ctx)
	// 根据操作类型获取对应的邮件通知模板
	notice, err = noticeDao.GetNoticByID(s.OperationType)
	if err != nil {
		// 获取通知模板失败
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}

	// 构建邮件发送地址（可能为绑定邮箱或重置链接）
	address = conf.ValidEmail + token

	// 替换邮件模板中的占位符 "Email" 为实际地址
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "email", address, -1)

	// 创建邮件对象
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail) // 设置发件人
	m.SetHeader("To", s.Email)          // 设置收件人
	m.SetHeader("Subject", "FanOne")    // 设置邮件主题
	m.SetBody("text/html", mailText)    // 设置邮件内容（HTML 格式）

	// 配置 SMTP 连接参数
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS // 强制使用 TLS 加密

	// 发送邮件
	err = d.DialAndSend(m)
	if err != nil {
		// 邮件发送失败
		code = 40001 // 自定义错误码：邮件发送失败
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}

	// 邮件发送成功
	return serializer.Response{
		Status: code,
		Msg:    e.GetMSG(code),
	}
}

// VerifyEmail 认证邮件
func (s *VerifyEmailService) VerifyEmail(ctx context.Context, token string) serializer.Response {
	var (
		userId        int
		email         string
		password      string
		operationType int
	)
	code := http.StatusOK
	// 验证token
	if token == "" {
		code = 30010 // token不存在
	} else {
		claims, err := mytools.ParseEmailToken(token)
		fmt.Println(claims.ID)
		if err != nil {
			code = 30007 // token认证失败
		} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
			code = 30008 // token过期
		} else {
			userId = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	if code != http.StatusOK {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
		}
	}

	// 通过token获取用户信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserByID(userId)
	if err != nil {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
		}
	}

	switch operationType {
	case 0:
		user.Email = email
	case 1:
		user.Email = ""
	case 2:
		err = user.SetPassword(password)
		if err != nil {
			code = http.StatusBadRequest
			return serializer.Response{
				Status: code,
				Msg:    e.GetMSG(code),
			}
		}
	}
	// 更新用户信息
	err = userDao.UpdateUserByID(userId, user)
	if err != nil {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMSG(code),
		Data:   serializer.BuildUser(user),
	}
}

// Register 用户注册
func (us *UserService) Register(context context.Context) serializer.Response {
	var user model.User
	code := http.StatusOK
	if us.Key == "" || len(us.Key) != 16 {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  "密钥为空或长度不足",
		}
	}

	// 加密操作
	mytools.Encrypt.SetKey(us.Key)

	userDao := dao.NewUserDao(context)
	_, exist, err := userDao.ExistOrNotByUserName(us.UserName)
	if err != nil {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
		}
	}
	if exist {
		code = 30001 // 用户名已存在
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
		}
	}
	var useMoney string
	// 初始金额加密
	if us.Money == "" {
		useMoney, err = mytools.Encrypt.MyEncrypt("0")
	} else {
		useMoney, err = mytools.Encrypt.MyEncrypt(us.Money)
		if err != nil {
			code = http.StatusBadRequest
			return serializer.Response{
				Status: code,
				Msg:    e.GetMSG(code),
			}
		}
	}

	user = model.User{
		UserName: us.UserName,
		NickName: us.NickName,
		Status:   model.Active,
		Avatar:   "defaultCat.jpg", // 注册用户默认头像
		Money:    useMoney,         // 使用加密后的金额
		Email:    us.Email,
	}
	// 密码加密
	if err3 := user.SetPassword(us.PassWord); err3 != nil {
		code = 30002 // 密码加密失败
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
		}
	}

	// 创建用户
	err4 := userDao.CreateUser(&user)
	if err4 != nil {
		code = http.StatusBadRequest
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMSG(code),
	}
}

// Login 用户登录
func (us *UserService) Login(ctx context.Context) serializer.Response {
	//var user *model.User
	code := http.StatusOK
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(us.UserName)
	// 判断用户是否存在
	if !exist || err != nil {
		code = 30003 // 用户不存在
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Data:   "用户不存在，请先注册",
		}
	}
	// 校验密码
	if !user.CheckPassword(us.PassWord) {
		code = 30004 // 密码错误
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Data:   "密码错误，请检查密码",
		}
	}
	// 添加token --> 因为http是无状态的
	token, err2 := mytools.CreateToken(int(user.ID), us.UserName, 0)
	if err2 != nil {
		code = 30005 // token下发失败
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Msg: e.GetMSG(code),
	}
}

// Update 用户数据更新
func (us *UserService) Update(ctx context.Context, userID int) serializer.Response {
	var user *model.User
	var err error
	var useMoney string
	code := http.StatusOK
	// 找到这个用户
	UserDao := dao.NewUserDao(ctx)
	user, err = UserDao.GetUserByID(userID) // 处理获取用户信息可能出现的错误
	if err != nil {
		code = 30003 // 用户不存在
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}
	// 昵称有改变-修改昵称
	if us.NickName != "" {
		user.NickName = us.NickName
	}

	if us.Money != "" {
		// 检查密钥是否有效
		if us.Key == "" || len(us.Key) != 16 {
			code = http.StatusBadRequest
			return serializer.Response{
				Status: code,
				Msg:    e.GetMSG(code),
				Error:  "密钥为空或长度不足",
			}
		}
		// 金额加密
		mytools.Encrypt.SetKey(us.Key)

		// 将新输入的金额转换为浮点数
		newMoney, err := strconv.ParseFloat(us.Money, 64)
		if err != nil {
			code = http.StatusBadRequest
			return serializer.Response{
				Status: code,
				Msg:    e.GetMSG(code),
				Error:  "输入的金额格式不正确",
			}
		}

		// 解密原有的余额
		oldMoneyStr, err := mytools.Encrypt.MyDecrypt(user.Money)
		if err != nil {
			code = http.StatusBadRequest
			return serializer.Response{
				Status: code,
				Msg:    e.GetMSG(code),
				Error:  "原余额解密失败",
			}
		}
		oldMoney, err := strconv.ParseFloat(oldMoneyStr, 64)
		if err != nil {
			code = http.StatusBadRequest
			return serializer.Response{
				Status: code,
				Msg:    e.GetMSG(code),
				Error:  "原余额转换失败",
			}
		}

		// 在原有余额基础上进行增减操作
		newBalance := oldMoney + newMoney

		// 加密新的余额
		useMoney, err = mytools.Encrypt.MyEncrypt(fmt.Sprintf("%.2f", newBalance))
		if err != nil {
			code = http.StatusBadRequest
			return serializer.Response{
				Status: code,
				Msg:    e.GetMSG(code),
				Error:  "新余额加密失败",
			}
		}
		user.Money = useMoney
	}

	err = UserDao.UpdateUserByID(userID, user)
	if err != nil {
		code = 30006 // 用户数据更新失败
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMSG(code),
		Data:   serializer.BuildUser(user),
	}
}

// PostFile 用户头像更新
func (us *UserService) PostFile(ctx context.Context, id int, file multipart.File, size int64) serializer.Response {
	// 上传到本地
	code := http.StatusOK
	var (
		user *model.User
		err  error
	)
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserByID(id)
	if err != nil {
		code = 30003 // 用户不存在
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}
	// 保存图片到本地
	path, err := UploadAvatarToLocalStatic(file, id, user.UserName)
	if err != nil {
		code = 30009
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}
	user.Avatar = path
	err = userDao.UpdateUserByID(id, user)
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
		Data:   serializer.BuildUser(user),
	}
}
