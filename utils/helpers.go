package utils

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// generateVerificationCode creates a random 6-digit code
func GenerateVerificationCode() string {
	rand.Seed(time.Now().UnixNano()) // Ensure randomness
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
