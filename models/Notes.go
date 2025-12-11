package models

import "gorm.io/gorm"

type Notes struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255);not null"`
	Content string `gorm:"type:text;not null"`
	UserID  uint
}
