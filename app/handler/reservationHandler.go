package handler

import (
	"fmt"
	"net/http"
	"time"

	// "github.com/Kantaro0829/go-gin-test/infra"
	"github.com/Kantaro0829/go-gin-test/json"
	"github.com/Kantaro0829/go-gin-test/model"
	"github.com/gin-gonic/gin"
)

// DBから取り出したデータを代入するstruct
type reseClass struct {
	RoomNo string
	Stime  string
	Etime  string
}

// 予約データを取得
// n号館n階の予約状況だけ返す
func ReservationInfo(c *gin.Context) {

	towerNumStr := c.Param("tower")
	tower := towerNumStr + "%"

	// 本日の日付を指定
	today := (time.Now().Format("2006-01-02"))

	// n号館n階の予約状況を変数に入れる
	// 予約状況を入れる変数
	class_rese := []model.Reservation{}
	result := db.Table("reservations").
		Select("room_no, s_time, e_time").
		Where("room_no LIKE ? AND rese_date LIKE ?", tower, today).
		Scan(&class_rese)

	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"status": 400})
		return
	}

	// 予約がない場合は0を返すようにする
	if len(class_rese) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "0"})
	} else {
		json := createReservationInfoJson(class_rese)
		c.JSON(http.StatusOK, json)
	}

}

// reservationsテーブルに格納されている予約をJson形式に書き換えている
func createReservationInfoJson(reseInfos []model.Reservation) []reseClass {
	//各教室の予約状況を格納するJson配列を作成
	reseInfo := []reseClass{}

	for _, v := range reseInfos {
		fmt.Printf("%v, %v, %v\n", v.RoomNo, v.STime, v.ETime)

		//各教室の予約状況を配列に格納する
		reseInfo = append(reseInfo, reseClass{
			RoomNo: v.RoomNo,
			Stime:  v.STime,
			Etime:  v.ETime,
		})
	}
	// ↓件数が増えたらいるかも
	fmt.Println("------------------出来上がったJson---------------------")
	fmt.Println(reseInfo)
	return reseInfo

}

//予約をする(insert)
func InsertReseInfo(c *gin.Context) {
	// 取得したjsonを格納する
	var reseJson json.JsonReservation

	if err := c.ShouldBindJSON(&reseJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 取得したJsonの中身を変数に格納する
	//teacherNo := reseJson.TeacherNo
	roomNo := reseJson.RoomNo
	reseDate := reseJson.ReseDate
	startTime := reseJson.StartT
	endTime := reseJson.EndT
	//stateNo := reseJson.StateNo
	fmt.Println(roomNo)
	fmt.Println(reseDate)
	fmt.Println(startTime)
	fmt.Println(endTime)

	reseI := []model.Reservation{}

	// ブッキングしてないか確かめる
	result := db.Table("reservations").
		Select("room_no, rese_date, s_time, e_time").
		Where("room_no LIKE ? AND rese_date = ?", roomNo, reseDate).
		Scan(&reseI)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "予約情報、取得失敗"})
		return
	}
	for _, v := range reseI {
		if (v.STime[:5] <= endTime && endTime < v.ETime) || (v.STime[:5] < startTime && endTime < v.ETime) || (v.STime[:5] < startTime && startTime < v.ETime) || (startTime < v.STime[:5] && v.ETime < endTime) {
			c.JSON(http.StatusOK, gin.H{"message": "0", "roomNO": v.RoomNo, "startTime": v.STime, "endTime": v.ETime})
			return
		}
	}

	// 予約表にデータを入れる
	// 今日の日付のフォーマット作成
	t := time.Now()
	const format = "2006-01-02"
	// データを作成
	resedata := model.Reservation{TeacherNo: 1, RoomNo: roomNo, ReseDate: reseDate, STime: startTime, ETime: endTime, Purpose: "面談", RequestDate: t.Format(format)}

	if err := db.Select("teacher_no", "room_no", "rese_date", "s_time", "e_time", "purpose", "request_date").Create(&resedata).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "1"})
	}
}
