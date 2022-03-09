package models

import "gorm.io/gorm"

type ChatData struct {
	gorm.Model
	TeacherId uint
	StudentId uint
	Message   string
}
