package db

import (
	"go-chat/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() error { //データベースのMigration
	return nil
}

func Create(teacherId uint, studentId uint, message string) error {
	dsn := "host=localhost user=admin password=admin port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.ChatData{})
	db.Create(&models.ChatData{TeacherId: teacherId, StudentId: studentId, Message: message})
	return nil
}
