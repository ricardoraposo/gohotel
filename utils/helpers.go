package utils

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func IsValidPassword(encryptedPassowrd, password string) bool {
    return bcrypt.CompareHashAndPassword([]byte(encryptedPassowrd), []byte(password)) == nil
}
