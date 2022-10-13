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

// var a = temp()

func GetRoomInfo(c *gin.Context) {

	roomNumStr := c.Param("roomNo")
	//fmt.Println(roomNumStr)
	roomNum, _ := strconv.ParseInt(roomNumStr, 10, 16)
	//roomNumberの上二桁だけ切り取り
	buildingNumAndFloor := roomNum / 100
	buildingAndFloor := strconv.FormatInt(buildingNumAndFloor, 10)
	buildingAndFloor = buildingAndFloor + "%"

	//db := infra.DBInit()

	today := time.Now()
	dayOfWeek := today.Weekday().String() // 曜日の取得
	fmt.Println(dayOfWeek[0:3]) //火曜なら "Tue" と表示される
	dayOfWeek = "Tue"

	rooms := []model.Room{}
	db.Order("room_no").
		Select("room_no").
		Where("room_no LIKE ?", buildingAndFloor).
		Find(&rooms)

	roomResults := []model.RoomResult{}
	roomScan := []model.RoomScan{}

	result := db.Order("timetables.room_no, timetables.time_no").Table("timetables").
		Select("timetables.room_no, timetables.time_no, timetables.subject_name").
		Where("timetables.room_no LIKE ? AND timetables.youbi = ?", buildingAndFloor, dayOfWeek).
		Scan(&roomResults)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "通常授業の教室使用情報、取得失敗"})
		return
	}

	//log
	fmt.Println("結合後のテーブル")
	for _, v := range roomResults {
		fmt.Println("-------------------------------------------------------------------")
		fmt.Printf("%v, %v, %v\n", v.RoomNo, v.TimeNo, v.SubjectName)
	}

	detectingResult := db.Order("room_no").Table("rooms").
		Select("room_no, is_detected").
		Where("room_no LIKE ?", buildingAndFloor).
		Scan(&roomScan)

	if detectingResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "センサーによる教室使用情報、取得失敗"})
		return
	}

	roomInfo := createRoomInfoJson(roomResults)
	reservationInfo := createReservationJson()
	detectingInfo := createDetectionJson(roomScan)
	response := AllInfo{NormalInfo: roomInfo, ReservationInfo: reservationInfo, DetectingInfo: detectingInfo}
	c.JSON(http.StatusOK, response)
}

type Class struct {
	TimeNo      string
	SubjectName string
}

type AllInfo struct {
	NormalInfo      map[string][]Class
	ReservationInfo map[string]string
	DetectingInfo   interface{}
}

func createReservationJson() map[string]string {
	reservationInfos := make(map[string]string)
	reservationInfos["reservation"] = "予約"
	return reservationInfos
}

func createDetectionJson(detectingInfo []model.RoomScan) interface{} {
	detections := make(map[string]bool)

	for _, v := range detectingInfo {
		fmt.Printf("%v,%v\n", v.RoomNo, v.IsDetected)
		detections[v.RoomNo] = v.IsDetected
	}
	fmt.Println("センサー情報")
	fmt.Println(detections)
	return detections
}

func createRoomInfoJson(roomInfos []model.RoomResult) map[string][]Class {
	//各教室の状況を格納するJson配列を作成する
	var currentRoomNo string                  //同じ教室番号を配列に分割するために判断する変数
	eachRoomInfos := make(map[string][]Class) //最終的に出力したいJsonの型宣言
	roomInfo := []Class{}

	for i, v := range roomInfos {
		fmt.Printf("%v, %v, %v\n", v.RoomNo, v.TimeNo, v.SubjectName)
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
	fmt.Println("------------------出来上がったJson---------------------")
	fmt.Println(eachRoomInfos)

	return eachRoomInfos

}

// func createRoomsSlice(rooms []model.Room) []string {
// 	roomSlice := []string{}
// 	for _, v := range rooms {
// 		roomSlice = append(roomSlice, v.RoomNo)
// 	}
// 	return roomSlice
// }
