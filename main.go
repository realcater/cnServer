package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("mysql", "root:asd890@tcp(127.0.0.1:3306)/chgk?charset=utf8&parseTime=True&loc=Local")
	db.Set("gorm:table_options", "charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{}, &Question{})

	r := gin.Default()
	r.POST("/login", Login)
	r.GET("/questions", GetAllQuestions)
	r.GET("/questions/:id", GetQuestion)
	r.POST("/questions", CreateQuestion)
	r.PUT("/questions/:id", UpdateQuestion)
	r.DELETE("/questions/:id", DeleteQuestion)

	r.RunTLS(":8080", "cert.pem", "key.pem")
}
