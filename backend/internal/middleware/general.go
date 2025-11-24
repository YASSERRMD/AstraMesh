package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/google/uuid"
)

func RequestIDMiddleware(c *fiber.Ctx) error {
	// Generate a new request ID if not present
	reqID := c.Get("X-Request-ID")
	if reqID == "" {
		reqID = uuid.New().String()
		c.Set("X-Request-ID", reqID)
	}
	c.Locals("request_id", reqID)
	return c.Next()
}

func RateLimiterMiddleware(max int) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        max,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).JSON(fiber.Map{
				"error": "Rate limit exceeded",
			})
		},
	})
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	// Default 500 status code
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Log the error (in production, use a proper logger)
	c.App().Config().ErrorHandler(c, err)

	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}