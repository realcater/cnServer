package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

//Question :
type Question struct {
	ID       uint    `json:"id"`
	Number   uint    `json:"number"`
	Question string  `json:"question"`
	Answer   string  `json:"answer"`
	Comments string  `json:"comments"`
	Source   string  `json:"source"`
	TourID   uint    `json:"tour_id"`
	AuthorID uint    `json:"author_id"`
	Rating   float64 `json:"rating"`
	TypeID   uint    `json:"type_id"`
}

// DeleteQuestion :
func DeleteQuestion(c *gin.Context) {
	if CheckToken(c) {
		id := c.Params.ByName("id")
		var question Question
		d := db.Where("id = ?", id).Delete(&question)
		fmt.Println(d)
		c.Status(http.StatusOK)
	}
}

// UpdateQuestion :
func UpdateQuestion(c *gin.Context) {
	if CheckToken(c) {
		var question Question
		id := c.Params.ByName("id")
		if err := db.Where("id = ?", id).First(&question).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		}
		c.BindJSON(&question)
		db.Save(&question)
		c.JSON(http.StatusCreated, question)
	}
}

//CreateQuestion :
func CreateQuestion(c *gin.Context) {
	if CheckToken(c) {
		var question Question
		c.BindJSON(&question)
		db.Create(&question)
		if db.NewRecord(question) {
			c.JSON(http.StatusBadRequest, gin.H{"message": createErrorMessage})
		} else {
			c.JSON(http.StatusCreated, question)
		}
	}
}

// GetQuestion :
func GetQuestion(c *gin.Context) {
	if CheckToken(c) {
		id := c.Params.ByName("id")
		var question Question
		if err := db.Where("id = ?", id).First(&question).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, question)
			fmt.Println(question)
		}
	}
}

// GetAllQuestions :
func GetAllQuestions(c *gin.Context) {
	if CheckToken(c) {
		var questions []Question
		if err := db.Find(&questions).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, questions)
		}
	}
}
