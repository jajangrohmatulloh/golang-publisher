package helper

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(formPassword string, dbPassword string) bool {
	// salt := "rahasia"
	// comparedPassword := formPassword + salt
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte("emptysecret"))
	if err != nil {
		fmt.Println("Password mismatch:", err)
        return false
	}
	fmt.Println("Password match!")
	return true
}
