package infra

import (
	"github.com/Kantaro0829/go-gin-test/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type SqlHandler struct {
// 	db *gorm.DB
// }

// func NewSqlHandler() database.SqlHandler {
// 	dsn := "root:ecc@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err.Error)
// 	}
// 	sqlHandler := new(SqlHandler)
// 	sqlHandler.db = db
// 	return sqlHandler
// }

func DBInit() *gorm.DB {
	dsn := "root:ecc@tcp(db:3306)/earlGrey?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.User{})
	// db.AutoMigrate(&model.Permmission{})
	// db.AutoMigrate((&model.Timer{}))

	return db
}
