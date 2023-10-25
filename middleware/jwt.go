package middleware

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

// JWT authentication using HS256 algorithm
func CreateToken(userId int, email string) (string, error) {

	// set claims (payload)
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// create token with claims (header)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// generate encoded token using the SECRET_KEY environment variable (signature)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}
