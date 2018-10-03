package main

import (
	"time"
)

const (
	expiredTokenMessage       = "Token is expired"
	invalidTokenMessage       = "Token is invalid"
	noTokenMessage            = "There is no token"
	createErrorMessage        = "Unable to create a question"
	tokenExpireTime           = time.Hour * 24
	tokenLength               = 64
	passwordSecurityIteration = 4096
	passwordSecurityKeyLen    = 64
	passwordSecuritySaltSize  = 32
)
