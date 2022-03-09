package db

import (
	"go-chat/server/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() error { //データベースのMigration
	return nil
}

func Create(user string, message string) error {
	dsn := "host=localhost user=admin password=admin port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.ChatHistory{})
	db.Create(&models.ChatHistory{User: user, Message: message})
	return nil
}
