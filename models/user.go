package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"username" form:"name"`
	Password string `gorm:"password" form:"password"`
	Email    string `gorm:"email" form:"email"`
}
