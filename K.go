package main

import (
	"time"
)

const (
	expiredTokenMessage = "Token is expired"
	invalidTokenMessage = "Token is invalid"
	noTokenMessage      = "There is no token"
	tokenExpireTime     = time.Hour * 24
	tokenLength         = 64
	createErrorMessage  = "Unable to create a question"
)
