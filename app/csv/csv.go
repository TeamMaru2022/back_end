package csv

import (
	"github.com/Kantaro0829/go-gin-test/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:ecc@tcp(db:3306)/maru?charset=utf8mb4&parseTime=True&loc=Local"
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

var cnt = 1

func Csv1(record [][]string) int {
	//(record[0][0])[3:] == "授業名"  はBOM付きCSVファイルの場合最初になにか入ってるからそれを省いて判定
	if !((record[0][0] == "授業名" || (record[0][0])[3:] == "授業名") && record[0][1] == "曜日" && record[0][2] == "授業時間" && record[0][3] == "教室番号") {
		return 1
	}

	// 前に入っている時間割を削除
	db.Exec("DELETE FROM timetables")

	// 0行目には "授業名" などの文字が入っているのでその行は飛ばす
	for i := 1; i < len(record); i++ {

		// 教室(RoomNo)が空欄の場合、その行は飛ばす
		// 何行飛ばしたかカウントする
		if len(record[i][3]) == 0 {
			cnt++
			continue
		} else {
			cnt = 1
		}

		// 空欄の場合は上の行の授業名を持ってくる
		// 授業名
		if len(record[i][0]) == 0 {
			record[i][0] = record[i-cnt][0]
		}
		// 曜日
		if len(record[i][1]) == 0 {
			record[i][1] = record[i-cnt][1]
		}
		// 授業時間
		if len(record[i][2]) == 0 {
			record[i][2] = record[i-cnt][2]
		}
		room := model.Timetable{SubjectName: (record[i][0]), Youbi: (record[i][1]), TimeNo: (record[i][2]), RoomNo: (record[i][3])}

		if err := db.Select("subject_name", "youbi", "time_no", "room_no").Create(&room).Error; err != nil {
			return 2
		}
	}
	return 0
}
