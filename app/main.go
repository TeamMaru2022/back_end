package main

import (
	"github.com/Kantaro0829/go-gin-test/handler"
	"github.com/Kantaro0829/go-gin-test/infra"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	infra.DBInit()
	router := gin.Default()

	// ここからCorsの設定
	router.Use(cors.Default())

	user := router.Group("/user")
	{
		user.GET("/get", handler.Getting)
		user.PUT("/reg", handler.UserReg)
		user.POST("/login", handler.UserLogin)
		user.PUT("/update", handler.UpdateUser)
		user.DELETE("/delete", handler.DeleteUser)
		//user.POST("/validate", handler.SampleJwtValidation)

	}
	reservation := router.Group("/reservation")
	{
		reservation.GET("/getRese", handler.ReservationInfo)
	}

	room := router.Group("/room")
	{
		room.GET("/:roomNo", handler.GetRoomInfo)
	}

	sensor := router.Group("/sensor")
	{
		sensor.POST("update", handler.UpdateDetectingInfo)
	}

	teacher := router.Group("/teacher")
	{
		teacher.POST("reg", handler.TeacherReg)
		teacher.POST("login", handler.TeacherLogin)
	}

	csv := router.Group("/csv")
	{
		csv.GET("csv", handler.ChangeCsv)
	}
	router.Run(":3000")

}
