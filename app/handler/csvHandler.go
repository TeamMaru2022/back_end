package handler

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"

	mycsv "github.com/Kantaro0829/go-gin-test/csv"

	"github.com/gin-gonic/gin"
)

func ChangeCsv(c *gin.Context) {

	// postで受け取ったファイルを読み込む
	file, _ := c.FormFile("timetable")

	// fileを開く
	file1, err := file.Open()
	if err != nil {
		log.Fatal(err)
	}

	// csvファイルを読み出す
	// records : [][]string
	reader := csv.NewReader(file1)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// DBにデータを入れる
	flag := mycsv.Csv1(records)

	// flagによって返すmessageを変更する
	// 0 : statusOK
	// 1 : csvファイルの形式が違う
	// 2 : DBに入れるときにエラーがでた
	fmt.Println(flag)
	if(flag == 0){
		c.JSON(http.StatusOK, gin.H{"message": "時間割の登録完了しました。"})
	}else if(flag == 1){
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "csvファイルの形式が違います。"})
	}else if(flag == 2){
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "時間割を登録できませんでした。"})
	}

}
