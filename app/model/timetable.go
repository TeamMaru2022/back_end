package model

type Timetable struct {
	No          uint16 `gorm:"primaryKey"`
	RoomNo      string
	SubjectName string
	Youbi       string
	TimeNo      string
}
