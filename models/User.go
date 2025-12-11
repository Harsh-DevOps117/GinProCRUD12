package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement;uniqueIndex"`
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Password string
	Notes    []Notes `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
