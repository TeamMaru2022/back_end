package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/Kantaro0829/go-gin-test/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:ecc@tcp(db:3306)/maru?charset=utf8mb4&parseTime=True&loc=Local"
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func Csv1() {
	csvFile, err := os.Open("csv/timetable.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db.Exec("DELETE FROM timetables")

	for i := 0; i < len(record); i++ {
		fmt.Print(record[i][0])

		room := model.Timetable{SubjectName: (record[i][0]), Youbi: (record[i][1]), TimeNo: (record[i][2]), RoomNo: (record[i][3])}

		if err := db.Select("subject_name", "youbi", "time_no", "room_no").Create(&room).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
}