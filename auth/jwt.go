package auth

import (
	"os"
	"time"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateJWT(userID uint) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	// Define token expiration time (e.g., 1 hour)
	expirationTime := time.Now().Add(1 * time.Hour)

	// Create the JWT claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		Issuer:    "task-app",
		Subject:   strconv.Itoa(int(userID)),
	}

	// Create JWT token with claims and secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
