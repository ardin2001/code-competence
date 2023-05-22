package middleware

import (
	"echo_golang/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func CreateToken(userId uint, name string) (string, error) {
	claims := &models.JwtCustomClaims{
		ID:   userId,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 100)),
		},
	}

	// Create token with claims
	godotenv.Load()
	dbHost := os.Getenv("SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(dbHost))

}
