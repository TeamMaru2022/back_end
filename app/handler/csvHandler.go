package handler

import (
	_ "fmt"

	"github.com/Kantaro0829/go-gin-test/csv"

	"github.com/gin-gonic/gin"
)

func ChangeCsv(c *gin.Context){
	csv.Csv1()
}