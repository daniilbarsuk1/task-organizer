package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"username" form:"username"`
	Password string `gorm:"password" form:"password"`
}
