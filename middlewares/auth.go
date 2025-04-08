package middlewares

import (
	"fmt"
	"os"
	"task/utils"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v2"

	"strings"
)

// AuthMiddleware is a middleware to protect routes requiring authentication
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the token from the "Authorization" header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Authorization header is missing", nil)
		}

		// Split the "Bearer <token>" format
		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid token format", nil)
		}

		claims := &jwt.RegisteredClaims{}

			token, err := jwt.ParseWithClaims(tokenString[1], claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if err != nil || !token.Valid {
				return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid or expired token", err.Error())
			}

			// Store user ID (from Subject) in context
			c.Locals("userID", claims.Subject)
			return c.Next()

	}
}

