package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

/*
type server struct {
	db     *gorm.DB
	routes *gin.Engine
}
*/
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

	checkTokenGroup := r.Group("/")
	checkTokenGroup.Use(CheckToken())
	{
		checkTokenGroup.GET("/questions", GetAllQuestions)
		checkTokenGroup.GET("/questions/:id", GetQuestion)
		checkTokenGroup.POST("/questions", CreateQuestion)
		checkTokenGroup.PUT("/questions/:id", UpdateQuestion)
		checkTokenGroup.DELETE("/questions/:id", DeleteQuestion)
	}
	r.RunTLS(":8080", "cert.pem", "key.pem")
}
