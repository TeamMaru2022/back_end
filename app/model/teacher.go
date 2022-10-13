package model

import (
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	ID          uint16 `gorm:"primaryKey auto_increment"`
	TeacherName string
	Mail        string `gorm:"unique"`
	PerNo       uint8
	Password    string `gorm:"not null"`
}
