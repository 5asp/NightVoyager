package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func CreatePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(old, new string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(old), []byte(new))
	if err != nil {
		return false
	} else {
		return true
	}
}
