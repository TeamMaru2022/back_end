package handler

import (
	"fmt"
	"net/http"

	"github.com/Kantaro0829/go-gin-test/model"
	"gorm.io/gorm"

	"github.com/Kantaro0829/go-gin-test/json"

	"github.com/gin-gonic/gin"
)

func UpdateDetectingInfo(c *gin.Context) {
	var sensorJson json.SensorInfoJson //受け取るJson配列の型宣言app/json/jsonRequest

	//上で宣言した構造体にJsonをバインド。エラーならエラー処理を返す
	if err := c.ShouldBindJSON(&sensorJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//それぞれJson配列の値を変数に代入
	roomNo := sensorJson.RoomNo
	isDetected := sensorJson.IsDetected

	fmt.Println("Jsonの値")
	fmt.Println(roomNo)
	fmt.Println(isDetected)

	//db := infra.DBInit()
	tx := db.Session(&gorm.Session{SkipDefaultTransaction: true})
	room := model.Room{}

	// DBに入っているデータを取り出す
	result := tx.Table("rooms").Select("is_detected").Where("room_no = ?", roomNo).Scan(&room)

	if result.Error != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": result.Error})
		return
	}
	fmt.Println("テーブルから取り出した検知したかの値")
	fmt.Println(room.IsDetected)

	// DBの値と教室の状態が同じかの判定
	if isDetected == room.IsDetected {
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{"status": 200, "message": "以前の検知結果と同じです。"})
		return
	}

	// DB値を更新
	if result = db.Model(&room).Where("room_no = ?", roomNo).Update("is_detected", isDetected); result.Error != nil {
		fmt.Println("データのの更新ができていません")
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": 503})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "登録完了"})
}
