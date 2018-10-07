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
func (s *Server) DeleteQuestion(c *gin.Context) {
	//	if checkToken(c) {
	id := c.Params.ByName("id")
	var question Question
	d := s.db.Where("id = ?", id).Delete(&question)
	fmt.Println(d)
	c.Status(http.StatusOK)
	//	}
}

// UpdateQuestion :
func (s *Server) UpdateQuestion(c *gin.Context) {
	//	if checkToken(c) {
	var question Question
	id := c.Params.ByName("id")
	if err := s.db.Where("id = ?", id).First(&question).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	}
	c.BindJSON(&question)
	s.db.Save(&question)
	c.JSON(http.StatusCreated, question)
	//	}
}

//CreateQuestion :
func (s *Server) CreateQuestion(c *gin.Context) {
	//	if checkToken(c) {
	var question Question
	c.BindJSON(&question)
	s.db.Create(&question)
	if s.db.NewRecord(question) {
		c.JSON(http.StatusBadRequest, gin.H{"message": createErrorMessage})
	} else {
		c.JSON(http.StatusCreated, question)
	}
	//	}
}

// GetQuestion :
func (s *Server) GetQuestion(c *gin.Context) {
	//	if checkToken(c) {
	id := c.Params.ByName("id")
	var question Question
	if err := s.db.Where("id = ?", id).First(&question).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, question)
		fmt.Println(question)
	}
	//	}
}

// GetAllQuestions :
func (s *Server) GetAllQuestions(c *gin.Context) {
	var allQuestions []Question
	if err := s.db.Find(&allQuestions).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, allQuestions)
	}
}
