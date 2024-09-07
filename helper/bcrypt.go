package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(formPassword string, dbPassword string) bool {
	salt := "rahasia"
	comparedPassword := formPassword + salt
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(comparedPassword))
	if err != nil {
        return false
	}
	return true
}
