package model

import (
	"gorm.io/gorm"
)

type Timetable struct {
	gorm.Model
	No          uint16 `gorm:"primaryKey"`
	RoomNo      uint16
	SubjectName string
	Youbi       string
	TeacherNo   uint16
	TimeNo      string
}
