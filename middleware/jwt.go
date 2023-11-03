package middleware

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	ID uint   `json:"id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// JWT authentication using HS256 algorithm
func CreateToken(userId int, email string) string {

	// set claims (payload)
	var claims jwtCustomClaims

	claims.ID = uint(userId)
	claims.Email = email
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))

	// create token with claims (header)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	_, err := os.Stat(".env")
    if err == nil {
        if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Error loading .env file")
			os.Exit(1)
		}
    }
	// generate encoded token using the SECRET_KEY environment variable (signature)
	validToken, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return validToken
}

func ExtractToken(e echo.Context) (uint, string) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		email := claims["email"].(string)
		return uint(userId), email
	}
	return 0, ""
}
