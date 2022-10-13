package model

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	RoomNo     string `gorm:"primaryKey"`
	Outlet     string
	Lan				 bool
	IsDetected bool
}

type RoomScan struct {
	RoomNo     string
	IsDetected bool
}
