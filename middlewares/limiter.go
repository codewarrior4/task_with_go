package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func RateLimiter(minutes int) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        minutes,                                // Maximum requests allowed
		Expiration: time.Duration(minutes) * time.Minute, // Dynamic expiration time
		LimitReached: func(c *fiber.Ctx) error {
			// Estimate remaining time (assuming each request resets expiration)
			retryAfter := int(time.Duration(minutes) * time.Minute / time.Second)

			// Set the Retry-After header
			c.Set("Retry-After", fmt.Sprintf("%d", retryAfter))

			// Return JSON response with retry time
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":       "Too many requests, please try again later.",
				"retry_after": fmt.Sprintf("%d seconds", retryAfter),
			})
		},
	})
}


