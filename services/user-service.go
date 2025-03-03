package services

import (
	"errors"
	"task-organizer/database"
	"task-organizer/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func New() Service {
	s := Service{}
	s.DB = database.OpenDB()
	return s
}

func (s Service) UserExists(username string) bool {
	var user models.User
	if err := s.DB.Where("name = ?", username).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func (s Service) CreateUser(user models.User) error {
	if s.UserExists(user.Name) {
		return errors.New("user exists")
	}
	bytePassword := []byte(user.Password)
	encodedPassword, err := bcrypt.GenerateFromPassword(bytePassword, 5)
	user.Password = string(encodedPassword)
	if err != nil {
		return err
	}
	s.DB.Table("users").Create(&user)
	return nil
}
