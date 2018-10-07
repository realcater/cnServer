package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// CheckToken :
func (s *Server) CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		token := c.GetHeader("Token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": noTokenMessage})
			return
		}
		if err := s.db.Where("token = ?", token).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": invalidTokenMessage})
			return
		}
		if user.UpdatedAt.Add(tokenExpireTime).Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": expiredTokenMessage})
			return
		}
		c.Next()
	}
}
