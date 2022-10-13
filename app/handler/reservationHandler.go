package handler

import (
	"fmt"
	"net/http"

	"github.com/Kantaro0829/go-gin-test/infra"
	"github.com/Kantaro0829/go-gin-test/json"
	"github.com/Kantaro0829/go-gin-test/model"
	"github.com/gin-gonic/gin"
)

// 予約データを取得
func ReservationInfo(c *gin.Context) {
	// TODO: /modelにreservationsから取ってくる型宣言をするgoファイルを作る

	// db.goからmysql内のDBにアクセス
	db := infra.DBInit()

	// TODO: 予約テーブルからデータを取得(教室番号、開始時間、終了時間)
	rese := []model.Reservation{}
	// state_noが1(承認済み)の予約のみ取り出す
	result := db.Select("room_no, s_time, e_time").
		Where("state_no = 1").
		First(&rese)
		// Scan(&rese)

	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"status": 400})
		return
	}

	json := createReservationInfoJson(rese)
	c.JSON(http.StatusOK, json)
}

type reseClass struct {
	RoomNo string
	Stime  string
	Etime  string
}

// reservationsテーブルに格納されている予約をJson形式に書き換えている
func createReservationInfoJson(reseInfos []model.Reservation) map[string][]reseClass {
	//各教室の予約状況を格納するJson配列を作成
	var roomNo string
	reseRoomInfo := make(map[string][]reseClass)
	reseInfo := []reseClass{}

	for i, v := range reseInfos {
		fmt.Printf("%v, %v, %v\n", v.RoomNo, v.STime, v.ETime)

		// ループの最初をroomNoに代入
		if i == 0 {
			roomNo = v.RoomNo
		}

		if roomNo != v.RoomNo {
			// 前の教室番号と違う教室番号の場合新しい連想配列を作成
			reseRoomInfo[roomNo] = reseInfo

			reseInfo = []reseClass{}
			roomNo = v.RoomNo
		}

		//各教室の予約状況を配列に格納する
		reseInfo = append(reseInfo, reseClass{
			RoomNo: v.RoomNo,
			Stime:  v.STime,
			Etime:  v.ETime,
		})
	}
	// ↓件数が増えたらいるかも
	reseRoomInfo[roomNo] = reseInfo
	fmt.Println("------------------出来上がったJson---------------------")
	fmt.Println(reseRoomInfo)
	return reseRoomInfo

}

//予約をする(insert)
func insertReseInfo(c *gin.Context) {
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

	// db.goからmysql内のDBにアクセス
	data := infra.DBInit()

	reseI := []model.Reservation{}

	// status = 1:承認済み、時間帯・教室番号がブッキングしてない
	// rese := data.Select("room_no, teacher_no, rese_date, s_tiem, e_time").
	// 	Where("state_no = 1 AND rese_date = ? AND room_no = ?", reseDate, roomNo).
	// 	Scan(&reseI)

	// ブッキングしないために、該当するデータの個数を取得する
	// 教室番号・日付・開始時間・終了時間がブッキングしてないか
	exi := data.Select("IF(EXISTS(Select * FROM reservations WHERE rese_date = ? AND room_no = ? AND s_time = ? AND e_time = ?), 1, 0)",
		reseDate, roomNo, startTime, endTime).Scan(&reseI)

	// 当てはまるデータがなかった場合（予約可能な状態）
	if exi == nil {
		// insertの処理を書く

	}
}

// TODO:insertに使う構造体を作成する
type insertRese struct {
}
