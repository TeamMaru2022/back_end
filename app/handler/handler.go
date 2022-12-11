package handler

import (
	"fmt"
	"net/http"

	//"github.com/Kantaro0829/go-gin-test/infra"
	"github.com/Kantaro0829/go-gin-test/json"
	"github.com/Kantaro0829/go-gin-test/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:ecc@tcp(db:3306)/maru?charset=utf8mb4&parseTime=True&loc=Local"
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})


func Getting(c *gin.Context) {
	//db := infra.DBInit()
	timers := []model.Timer{}

	if result := db.Find(&timers); result.Error != nil {
		fmt.Println("データ取得失敗")
	}
	// fmt.Println(users)
	for i, v := range timers {
		fmt.Printf("%v回目", i)
		fmt.Println(v.TimeNo)
		fmt.Println(v.STime)
		fmt.Println(v.ETime)
		fmt.Println(v.ID)
	}

	c.JSON(http.StatusOK, gin.H{"message": "hello world"})
}

func UserReg(c *gin.Context) {
	var userJson json.JsonRequestUser //受け取るJson配列の型宣言app/json/jsonRequest

	//上で宣言した構造体にJsonをバインド。エラーならエラー処理を返す
	if err := c.ShouldBindJSON(&userJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//それぞれJson配列の値を変数に代入
	mail := userJson.UserMail
	age := userJson.UserAge
	password := userJson.UserPassword

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12) //パスワードハッシュ化

	if err != nil {
		panic("failed to hash password")
	}

	//db := infra.DBInit()
	user := model.User{Mail: mail, Password: hashedPassword, Age: age}

	if result := db.Create(&user); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"status": 400})
		return
	}
	fmt.Println("登録されたパスワード")
	fmt.Println(user.Password)
	//userのIdとメールアドレスを元にJWTを発行しているけどとりあえずここは無視してOK
	// 	token := auth.CreateTokenString(user.ID, user.Mail)

	// 	c.JSON(http.StatusOK, gin.H{"message": "data was inserted", "token": token})
	// }
}
func UserLogin(c *gin.Context) {
	var userLoginJson json.JsonRequestUser

	if err := c.ShouldBindJSON(&userLoginJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mail := userLoginJson.UserMail
	password := userLoginJson.UserPassword

	//db := infra.DBInit()
	user := model.User{}
	//select するときはこんな感じselect paasword, mail, id from users where mail = "test@gmail.com";
	result := db.Select("password", "mail", "id").Where("mail = ?", mail).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"status": 400})
		return
	}

	//パスワードがあっているかの確認
	if isAuthorized := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); isAuthorized != nil {

		fmt.Println("不一致")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "failed to login"})
		return

	}

	// token := auth.CreateTokenString(user.ID, user.Mail)
	// c.JSON(http.StatusOK, gin.H{"message": "succeed login", "token": token})

}

func UpdateUser(c *gin.Context) {
	var updateUserJson json.UpdateUserJson

	if err := c.ShouldBindJSON(&updateUserJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mail := updateUserJson.Mail
	password := updateUserJson.Password
	newMail := updateUserJson.NewMail
	newPassword := updateUserJson.NewPassword

	//db := infra.DBInit()
	user := model.User{}

	result := db.Select("password").Where("mail = ?", mail).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"status": 400})
		return
	}

	if isAuthorized := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); isAuthorized != nil {
		fmt.Println("不一致")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "パスワードかメールアドレスが正しくありません"})
		return
	}

	//新しいパスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), 12)
	if err != nil {
		panic("failed to hash password")
	}

	//dbに値を更新する
	db.First(&user)
	user.Mail = newMail
	user.Password = hashedPassword

	if result = db.Save(&user); result.Error != nil {
		fmt.Println("DBの更新ができていません")
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": 503})
		return
	}

	fmt.Println(result.RowsAffected)
	amountOfUpdate := result.RowsAffected

	if amountOfUpdate != 1 {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": 503})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 200})
}

func DeleteUser(c *gin.Context) {
	var deleteUserJson json.DeleteUserJson

	if err := c.ShouldBindJSON(&deleteUserJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mail := deleteUserJson.Mail
	password := deleteUserJson.Password

	//db := infra.DBInit()
	user := model.User{}

	result := db.Select("password", "ID").Where("mail = ?", mail).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"status": 400})
		return
	}

	if isAuthorized := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); isAuthorized != nil {

		fmt.Println("不一致")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "パスワードかメールアドレスが正しくありません"})
		return
	}

	if result = db.Where("mail = ?", mail).Delete(&user); result.Error != nil {
		fmt.Println("データ削除失敗")
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "正しくデータの削除を行えませんでした"})
		return
	}

	if result.RowsAffected != 1 {
		fmt.Println("データ削除失敗")
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "正しくデータの削除を行えませんでした"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "正常に削除できました"})

}

// func SampleJwtValidation(c *gin.Context) {
// 	var sampleValidationJson json.SampleValidationJson

// 	if err := c.ShouldBindJSON(&sampleValidationJson); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	token := sampleValidationJson.Token

// 	sample := auth.ValidateTokenString(token)
// 	c.JSON(http.StatusOK, gin.H{"mail": sample.Mail, "id": sample.Id})
// }
