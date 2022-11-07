package csv

import (
	"fmt"

	"github.com/Kantaro0829/go-gin-test/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:ecc@tcp(db:3306)/maru?charset=utf8mb4&parseTime=True&loc=Local"
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func Csv1(record [][]string) int {
	//(record[0][0])[3:] == "授業名"  はBOM付きCSVファイルの場合最初になにか入ってるからそれを省いて判定
	if !((record[0][0] == "授業名" || (record[0][0])[3:] == "授業名" ) && record[0][1] == "曜日" && record[0][2] == "授業時間" && record[0][3] == "教室番号") {
		return 1
	}

	db.Exec("DELETE FROM timetables")

	for i := 1; i < len(record); i++ {

		room := model.Timetable{SubjectName: (record[i][0]), Youbi: (record[i][1]), TimeNo: (record[i][2]), RoomNo: (record[i][3])}

		if err := db.Select("subject_name", "youbi", "time_no", "room_no").Create(&room).Error; err != nil {
			return 2
		}
	}
	return 0
}
