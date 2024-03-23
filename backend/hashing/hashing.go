package hashing

import (
	"log"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
