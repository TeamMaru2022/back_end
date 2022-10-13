package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	gorm.Model
	Mail          string `gorm:"unique"`
	Password      []byte
	Age           uint8
	SamapleSample string
}

type Timer struct {
	gorm.Model
	TimeNo string `gorm:"primaryKey"`
	STime  string `gorm:"not null"`
	ETime  string `gorm:"not null"`
}
