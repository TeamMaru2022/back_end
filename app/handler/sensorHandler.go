package handler

import (
	"fmt"
	"net/http"

	"github.com/Kantaro0829/go-gin-test/model"
	"gorm.io/gorm"

	"github.com/Kantaro0829/go-gin-test/json"

	"github.com/gin-gonic/gin"

	"log"

	"github.com/line/line-bot-sdk-go/linebot"
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
	// room.IsDetected: true , isDetected: true  の場合 line_beacon, is_detected をfalse
	// room.IsDetected: false, isDetected: true  の場合 is_detected をtrue
	// room.IsDetected: true , isDetected: false の場合変更なし
	if room.IsDetected && isDetected {
		// DB値を変更
		if result = db.Model(&room).Where("room_no = ?", roomNo).Updates(map[string]interface{}{"line_beacon": false, "is_detected": false}); result.Error != nil {
			fmt.Println("LineBeacon, センサーのデータ更新ができていません")
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": 503})
			return
		}
	} else if !room.IsDetected && isDetected {
		// DB値を変更
		if result = db.Model(&room).Where("room_no = ?", roomNo).Update("is_detected", true); result.Error != nil {
			fmt.Println("LineBeaconのデータ更新ができていません")
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": 503})
			return
		}
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "登録完了"})
}

func LineBeacon(c *gin.Context) {

	//HWIDに対応した教室番号のmap
	var roomno = map[string]string{"01679e6d47": "1201"}
	//ビーコンのHWID
	beaconhwid := "0"

	bot, err := linebot.New(
		//チャンネルシークレット
		"d9249c1e120f87d0b988e2c71f7042ad",
		//チャンネルアクセストークン
		"DpzLfnz4adplz7AGd4SPp7aHlDBcMdqX1LQ+WPN0geuNdth/oxDT21gFuTFdTcmmuCrcaqK5VQnvAOEuR4B4kYUAW4XHnTQN4e9M2aKuTADXtkGY68y2NLXNdyGMIOngP6IJzWoRxr0VxmokIZqqggdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(c.Request)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.Writer.WriteHeader(400)
		} else {
			c.Writer.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeBeacon {
			switch event.Type {
			case linebot.EventTypeBeacon:
				log.Println(" Beacon event....")
				//ビーコンイベントの判定
				if b := event.Beacon; b != nil {
					//ビーコンHWIDの取り出し
					beaconhwid = b.Hwid
					log.Print(beaconhwid)
					//ビーコンHWIDに対応した教室番号の判定
					if _, ok := roomno[beaconhwid]; ok {

						// lineBeaconでtrueが送られてきたときにDBに入れる
						room := model.Room{}
						// DBに入っているデータを取り出す
						tx := db.Session(&gorm.Session{SkipDefaultTransaction: true})
						result := tx.Table("rooms").Select("line_beacon").Where("room_no = ?", roomno[beaconhwid]).Scan(&room)
						if result.Error != nil {
							tx.Rollback()
							c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": result.Error})
							return
						}
						// DB値を更新
						if result = db.Model(&room).Where("room_no = ?", roomno[beaconhwid]).Update("line_beacon", true); result.Error != nil {
							fmt.Println("データの更新ができていません")
							c.JSON(http.StatusServiceUnavailable, gin.H{"status": 503})
							return
						}

						tx.Commit()
						c.JSON(http.StatusOK, gin.H{"message": "登録完了"})
					}
				}
			}
		}
	}
}
