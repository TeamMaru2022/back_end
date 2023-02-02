package handler

import (
	"fmt"
	"net/http"

	// "reflect"
	"time"

	//"github.com/Kantaro0829/go-gin-test/auth"
	//"github.com/Kantaro0829/go-gin-test/infra"
	"github.com/Kantaro0829/go-gin-test/model"

	//"github.com/Kantaro0829/go-gin-test/json"
	//"github.com/Kantaro0829/go-gin-test/model"
	//"github.com/gin-gonic/gin"
	//"golang.org/x/crypto/bcrypt"

	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRoomInfo(c *gin.Context) {

	roomNumStr := c.Param("roomNo")
	roomNum, _ := strconv.ParseInt(roomNumStr, 10, 16)
	// roomNumberの上二桁だけ切り取り
	buildingNumAndFloor := roomNum / 100
	buildingAndFloor := strconv.FormatInt(buildingNumAndFloor, 10)
	buildingAndFloor = buildingAndFloor + "%"

	today := (time.Now().Format("2006-01-02")) //本日の日付を取得
	dayOfWeek := time.Now().Weekday().String() // 曜日の取得
	fmt.Println(dayOfWeek)
	dayOfWeek = dayOfWeek[0:3] //火曜なら "Tue"

	fmt.Println(dayOfWeek)

	// 今何限目を取得
	timer := []model.Timer{}
	time := db.Order("time_no").
		Select("time_no, s_time, e_time").
		Table("timers").
		Scan(&timer)

	if time.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "時間テーブル、取得失敗"})
		return
	}

	// 通常授業の情報を取得
	roomResults := []model.RoomResult{}
	result := db.Order("timetables.room_no, timetables.time_no").Table("timetables").
		Select("timetables.room_no, timetables.time_no, timetables.subject_name").
		Where("timetables.room_no LIKE ? AND timetables.youbi = ?", buildingAndFloor, dayOfWeek).
		Scan(&roomResults)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "通常授業の教室使用情報、取得失敗"})
		return
	}

	// センサー情報取得
	roomScan := []model.RoomScan{}
	detectingResult := db.Order("room_no").Table("rooms").
		Select("room_no, is_detected").
		Where("room_no LIKE ?", buildingAndFloor).
		Scan(&roomScan)

	if detectingResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "センサーによる教室使用情報、取得失敗"})
		return
	}

	// 予約状況を取得
	class_rese := []model.Reservation{}
	resersult := db.Order("room_no").Table("reservations").
		Select("room_no, s_time, e_time").
		Where("room_no LIKE ? AND rese_date LIKE ?", buildingAndFloor, today).
		Scan(&class_rese)

	if resersult.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"status": 400})
		return
	}

	timerInfo := createTimerInfoJson(timer)
	roomInfo := createRoomInfoJson(roomResults)
	reservationInfo := createReservationJson(class_rese)
	detectingInfo := createDetectionJson(roomScan)
	response := AllInfo{TimerInfo: timerInfo, NormalInfo: roomInfo, ReservationInfo: reservationInfo, DetectingInfo: detectingInfo}
	c.JSON(http.StatusOK, response)
}

type Class struct {
	TimeNo      string
	SubjectName string
}

type Rese struct {
	STime string
	ETime string
}

type Timer struct {
	TimeNo int
	STIME  string
	ETIME  string
}

type AllInfo struct {
	TimerInfo       int
	NormalInfo      map[string][]Class
	ReservationInfo map[string][]Rese
	DetectingInfo   interface{}
}

func createDetectionJson(detectingInfo []model.RoomScan) interface{} {
	detections := make(map[string]bool)

	for _, v := range detectingInfo {
		if v.IsDetected && v.LineBeacon {
			// 人感センサーとBeaconが true なら教室に誰かいる
			detections[v.RoomNo] = true
		} else {
			// 人感センサーかBeaconが false なら教室に誰もいない
			detections[v.RoomNo] = false
		}
	}
	fmt.Println("\n\n-----------センサー情報-----------")
	fmt.Println(detections)
	fmt.Println()
	fmt.Println()
	return detections
}

// 教室の予約状況を送る
func createReservationJson(reservationJsons []model.Reservation) map[string][]Rese {
	var currentRoomNo string
	eachReservationInfo := make(map[string][]Rese)
	reservationJson := []Rese{}

	for i, v := range reservationJsons {
		if i == 0 {
			currentRoomNo = v.RoomNo
		}

		if currentRoomNo != v.RoomNo {
			//以前の教室番号と違う教室番号の場合新しい連想配列を作る
			eachReservationInfo[currentRoomNo] = reservationJson
			//各教室1~5限情報をを格納する配列の初期化
			reservationJson = []Rese{}
			currentRoomNo = v.RoomNo
		}

		//各教室の1〜5限目の情報を格納する配列に値を入れる
		reservationJson = append(reservationJson, Rese{
			STime: v.STime,
			ETime: v.ETime,
		})
	}

	//最後だけfor文が回らないので
	eachReservationInfo[currentRoomNo] = reservationJson

	return eachReservationInfo
}

// 教室の授業予定を送る
func createRoomInfoJson(roomInfos []model.RoomResult) map[string][]Class {
	//各教室の状況を格納するJson配列を作成する
	var currentRoomNo string                  //同じ教室番号を配列に分割するために判断する変数
	eachRoomInfos := make(map[string][]Class) //最終的に出力したいJsonの型宣言
	roomInfo := []Class{}

	for i, v := range roomInfos {
		if i == 0 {
			//ループの最初は変数currentRoomに代入
			currentRoomNo = v.RoomNo
		}

		if currentRoomNo != v.RoomNo {
			//以前の教室番号と違う教室番号の場合新しい連想配列を作る
			eachRoomInfos[currentRoomNo] = roomInfo
			//各教室1~5限情報をを格納する配列の初期化
			roomInfo = []Class{}
			currentRoomNo = v.RoomNo
		}

		//各教室の1〜5限目の情報を格納する配列に値を入れる
		roomInfo = append(roomInfo, Class{
			TimeNo:      v.TimeNo,
			SubjectName: v.SubjectName,
		})

	}
	//最後だけfor文が回らないので
	eachRoomInfos[currentRoomNo] = roomInfo
	fmt.Println("\n\n------------------出来上がったJson---------------------")
	fmt.Println(eachRoomInfos)

	return eachRoomInfos

}

func createTimerInfoJson(timerInfos []model.Timer) int {
	// 今何限目かを返す関数

	// 現在時刻 : string
	const TimeFormat = "15:04:05"
	nowTime := time.Now().Format(TimeFormat)

	for _, v := range timerInfos {

		if v.STime < nowTime && nowTime < v.ETime {
			fmt.Println(v.TimeNo)
			return v.TimeNo
		}
	}
	return 0
}
