package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserName string
	PassWord string
	Avatar   string
}
