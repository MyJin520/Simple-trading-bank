package model

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	UserID  int    `gorm:"not null"`
	Name    string `gorm:"type:varchar(10) not null"`
	Phone   string `gorm:"type:varchar(11) not null"`
	Address string `gorm:"varchar(50) not null"`
}
