package utils

import (
	"expense-tracker/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(userPassword string, providedPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))

	if err != nil {
		logger.Error("password did not match")
		return false
	}
	return true
}
