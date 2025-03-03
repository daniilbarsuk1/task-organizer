package database

import (
	"task-organizer/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB() *gorm.DB {
	dsn := "user=postgres password=admin dbname=taskmanager port=5432 sslmode=disable TimeZone=Europe/Minsk"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	return db
}
