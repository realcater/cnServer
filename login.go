package main

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

// User :
type User struct {
	Username    string `gorm:"primary_key"`
	HashAndSalt string
	Token       string `gorm:"unique_index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// LoginData :
type LoginData struct {
	Username string
	Password string
}

// APIAnswer :
type APIAnswer struct {
	Username   string
	Authorized bool
	UserIsNew  bool
	NewToken   string
}

// Login :
func (s *Server) Login(c *gin.Context) {
	var user User
	var loginData LoginData
	var answer APIAnswer
	var httpStatus int
	c.BindJSON(&loginData)
	if err := s.db.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		httpStatus, answer = s.createNewUser(loginData)
	} else {
		httpStatus, answer = s.checkUser(&user, loginData)
	}
	c.JSON(httpStatus, answer)
}

func (s *Server) checkUser(user *User, loginData LoginData) (httpStatus int, answer APIAnswer) {
	if checkHash(loginData.Password, user.HashAndSalt) {
		s.giveUserAToken(user)
		answer = APIAnswer{user.Username, true, false, user.Token}
		httpStatus = http.StatusOK
	} else {
		answer = APIAnswer{user.Username, false, false, ""}
		httpStatus = http.StatusUnauthorized
	}
	return
}

func (s *Server) createNewUser(loginData LoginData) (httpStatus int, answer APIAnswer) {
	var user User
	user.Username = loginData.Username
	user.HashAndSalt = makeHashAndSalt(loginData.Password)
	s.db.Create(&user)
	s.giveUserAToken(&user)
	answer = APIAnswer{user.Username, true, true, user.Token}
	httpStatus = http.StatusCreated
	return
}

func (s *Server) giveUserAToken(user *User) {
	user.Token = base64.StdEncoding.EncodeToString(randomBytes(tokenLength))
	s.db.Save(user)
}
