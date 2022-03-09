package models

import "gorm.io/gorm"

type ChatHistory struct {
	gorm.Model
	User    string
	Message string
}
