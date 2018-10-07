package main

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/subtle"
	"encoding/base64"
	"golang.org/x/crypto/pbkdf2"
)

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
