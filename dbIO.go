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
	id := c.Params.ByName("id")
	var person Question
	d := db.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(http.StatusOK, gin.H{"id #" + id: "deleted"})
}

// UpdateQuestion :
func UpdateQuestion(c *gin.Context) {
	var person Question
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	}
	c.BindJSON(&person)
	db.Save(&person)
	c.JSON(http.StatusOK, person)
}

//CreateQuestion :
func CreateQuestion(c *gin.Context) {
	var person Question
	c.BindJSON(&person)
	db.Create(&person)
	c.JSON(http.StatusOK, person)
}

// GetQuestion :
func GetQuestion(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Question
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, person)
		fmt.Println(person)
	}
}

// GetQuestions :
func GetQuestions(c *gin.Context) {
	var question []Question
	if err := db.Find(&question).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, question)
	}
}