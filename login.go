package main

// only need mysql OR sqlite
// both are included here for reference
import (
	//"fmt"
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"time"
)

// User :
type User struct {
	Username     string `gorm:"primary_key"`
	HashPassword string
	Token        string `gorm:"unique_index"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
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
func Login(c *gin.Context) {
	var user User
	var loginData LoginData
	c.BindJSON(&loginData)
	if err := db.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		CreateNewUser(c, loginData)
	} else {
		CheckUser(c, &user, loginData)
	}
}

// CheckUser :
func CheckUser(c *gin.Context, user *User, loginData LoginData) {
	if user.HashPassword == loginData.Password {
		GiveUserAToken(c, user)
		answer := APIAnswer{user.Username, true, false, user.Token}
		c.JSON(http.StatusOK, answer)
	} else {
		answer := APIAnswer{user.Username, false, false, ""}
		c.JSON(http.StatusUnauthorized, answer)
	}
}

// CreateNewUser :
func CreateNewUser(c *gin.Context, loginData LoginData) {
	var user User
	user.Username = loginData.Username
	user.HashPassword = loginData.Password
	db.Create(&user)

	GiveUserAToken(c, &user)
	answer := APIAnswer{user.Username, true, true, user.Token}
	c.JSON(http.StatusOK, answer)
}

// CheckToken :
func CheckToken(c *gin.Context) bool {
	var user User
	token := c.GetHeader("Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": noTokenMessage})
		return false
	}
	if err := db.Where("token = ?", token).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": invalidTokenMessage})
		return false
	}
	if user.UpdatedAt.Add(tokenExpireTime).Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": expiredTokenMessage})
		return false
	}
	return true
}

// GiveUserAToken :
func GiveUserAToken(c *gin.Context, user *User) {
	user.Token = GenerateToken()
	db.Save(user)
}

// GenerateToken :
func GenerateToken() string {
	randomBytes := make([]byte, 12)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(randomBytes)
}
