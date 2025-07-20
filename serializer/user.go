package serializer

import (
	"go-store/conf"
	"go-store/model"
)

type User struct {
	// VO view Objective-->传给前端的数据类型
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"createAt"`
}

func BuildUser(user *model.User) *User {
	return &User{
		ID:       int(user.ID),
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   conf.Host + conf.HttpPort + conf.AvatarPath + user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
	}
}
