package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Server :
type Server struct {
	db    *gorm.DB
	route *gin.Engine
}

func main() {
	db, err := gorm.Open("mysql", "root:asd890@tcp(127.0.0.1:3306)/chgk?charset=utf8&parseTime=True&loc=Local")
	db.Set("gorm:table_options", "charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{}, &Question{})

	s := Server{db, gin.Default()}
	s.route.POST("/login", s.Login)

	checkTokenGroup := s.route.Group("/")
	checkTokenGroup.Use(s.CheckToken())
	{
		checkTokenGroup.GET("/questions", s.GetAllQuestions)
		checkTokenGroup.GET("/questions/:id", s.GetQuestion)
		checkTokenGroup.POST("/questions", s.CreateQuestion)
		checkTokenGroup.PUT("/questions/:id", s.UpdateQuestion)
		checkTokenGroup.DELETE("/questions/:id", s.DeleteQuestion)
	}
	s.route.RunTLS(":8080", "cert.pem", "key.pem")
}
