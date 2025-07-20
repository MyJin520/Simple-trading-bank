package serializer

import (
	"go-store/model"
	"go-store/pkg/mytools"
)

type Money struct {
	UserID    int    `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}

func BuildUserMoney(user *model.User, money string, key string) Money {
	mytools.Encrypt.SetKey(key)
	money2, _ := mytools.Encrypt.MyDecrypt(money)

	return Money{
		UserID:    int(user.ID),
		UserName:  user.UserName,
		UserMoney: money2,
	}
}
