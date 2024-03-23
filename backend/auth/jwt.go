package auth

import (
	"log"
	"os"
	"time"
	"buzzer/hashing"
	"buzzer/database"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func Create_JWT_Token(db *gorm.DB, username string, password string) (string, error) {
	// check if user exists with gorm
	admin, err := database.GetAdmin(db, username)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	
	// check if password is correct
	if !hashing.VerifyPassword(password, admin.Password) {
		log.Printf("Password for user %s is incorrect", username)
		return "", nil
	}
	
	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
}

func Verify_JWT_Token(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Println(err)
		return false, err
	}

	return token.Valid, nil
}