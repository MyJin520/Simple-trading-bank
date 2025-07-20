package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// 使用 uniqueIndex 标签确保 UserName 和 Email 组合唯一
	UserName string `gorm:"uniqueIndex:idx_user_name_email"`
	Email    string `gorm:"uniqueIndex:idx_user_name_email"`
	Password string
	NickName string // 昵称
	Status   string // 用户状态
	Avatar   string // 头像
	Money    string // 密文存储
}

const (
	PasswordEncryptionDifficulty int    = 12       // 密码加密难度
	Active                       string = "active" // 激活用户
)

// SetPassword 密码加密
func (user *User) SetPassword(pwd string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), PasswordEncryptionDifficulty)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 密码解密
func (user *User) CheckPassword(pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd))
	return err == nil
}
