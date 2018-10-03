package main

// only need mysql OR sqlite
// both are included here for reference
import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/subtle"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/pbkdf2"
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
func Login(c *gin.Context) {
	var user User
	var loginData LoginData
	c.BindJSON(&loginData)
	if err := db.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		createNewUser(c, loginData)
	} else {
		checkUser(c, &user, loginData)
	}
}

func checkUser(c *gin.Context, user *User, loginData LoginData) {
	if checkHash(loginData.Password, user.HashAndSalt) {
		giveUserAToken(c, user)
		answer := APIAnswer{user.Username, true, false, user.Token}
		c.JSON(http.StatusOK, answer)
	} else {
		answer := APIAnswer{user.Username, false, false, ""}
		c.JSON(http.StatusUnauthorized, answer)
	}
}

func createNewUser(c *gin.Context, loginData LoginData) {
	var user User
	user.Username = loginData.Username
	user.HashAndSalt = makeHashAndSalt(loginData.Password)
	db.Create(&user)

	giveUserAToken(c, &user)
	answer := APIAnswer{user.Username, true, true, user.Token}
	c.JSON(http.StatusOK, answer)
}

func giveUserAToken(c *gin.Context, user *User) {
	user.Token = base64.StdEncoding.EncodeToString(randomBytes(tokenLength))
	db.Save(user)
}

func makeHashAndSalt(password string) string {
	salt := randomBytes(passwordSecuritySaltSize)
	hash := makeHash(password, salt)
	return base64.StdEncoding.EncodeToString(append(salt, hash...))
}

func checkHash(password, hashAndSalt string) bool {
	bytesHashAndSalt, _ := base64.StdEncoding.DecodeString(hashAndSalt)
	salt := bytesHashAndSalt[:passwordSecuritySaltSize]
	hash := bytesHashAndSalt[passwordSecuritySaltSize:]
	return subtle.ConstantTimeCompare(makeHash(password, salt), hash) == 1
}

func makeHash(password string, salt []byte) []byte {
	return pbkdf2.Key([]byte(password), salt, passwordSecurityIteration, passwordSecurityKeyLen, sha1.New)
}

func randomBytes(len int) []byte {
	bytes := make([]byte, len)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return bytes
}
