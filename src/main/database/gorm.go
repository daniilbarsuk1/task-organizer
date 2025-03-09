package database

import (
	"log"
	"task-organizer/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB() (db *gorm.DB) {
	var err error
	dsn := "host=database user=postgres password=postgres dbname=taskmanager port=5432 sslmode=disable TimeZone=Europe/Minsk"
	for db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil; {
		log.Print(err)
	}
	db.AutoMigrate(&models.User{})
	return db
}
