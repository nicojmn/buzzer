package auth

import (
	"buzzer/database"
	"buzzer/hashing"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Create_JWT_Token(username string, password string) (string, error) {
	// check if user exists with gorm
	admin, err := database.GetAdmin(username)
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
		"exp": time.Now().Add(24 * time.Hour).Unix(),
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

func IsLogged(ctx *fiber.Ctx) bool {
	token := ctx.Cookies("jwt")
	if (token == "") {
		return false
	} else {
		verify, err := Verify_JWT_Token(token)
		if (err != nil) {
			return false
		}
		if (!verify) {
			return false
		}
	
		// retrieve username from token
		// check if user exists with gorm
		decoded, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			log.Println(err)
			return false
		}
		claims, ok := decoded.Claims.(*jwt.MapClaims)
		if !ok {
			log.Println("Error decoding claims")
			return false
		}

		if (*claims)["username"] == nil {
			return false
		}

		_, err = database.GetAdmin((*claims)["username"].(string))
		if (err != nil) {
			log.Println(err.Error())
			return false
		}


		return true
	}
}

func IsTeam(c *fiber.Ctx) bool {
	token := c.Cookies("jwt")
	if token == "" {
		return false
	} else {
		verify, err := Verify_JWT_Token(token)
		if err != nil {
			return false
		}
		if !verify {
			return false
		}
		// retrieve username from token
		// check if user exists with gorm
		decoded, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			log.Println(err)
			return false
		}
		claims, ok := decoded.Claims.(*jwt.MapClaims)
		if !ok {
			log.Println("Error decoding claims")
			return false
		}

		if (*claims)["teamName"] == nil {
			return false
		}

		_, err = database.GetTeam((*claims)["teamName"].(string))
		if err != nil {
			log.Println(err.Error())
			return false
		}
		return true
	}
}

func Create_Team_JWT_Token(teamName string, id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"teamName": teamName,
		"teamID": id,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println(err)
		return "", err
	}
	
	return tokenString, nil
}

func Retrieve_ID_from_JWT(token string) (int, error) {
	ok, err := Verify_JWT_Token(token)
	if err != nil || !ok {
		return -1, err
	}
	
	decoded, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Println(err)
		return -1, err
	}
	claims, ok := decoded.Claims.(*jwt.MapClaims)
	if !ok {
		log.Println("Error decoding claims")
		return -1, nil
	}

	if (*claims)["teamID"] == nil {
		return -1, nil
	}

	return int((*claims)["teamID"].(float64)), nil
}