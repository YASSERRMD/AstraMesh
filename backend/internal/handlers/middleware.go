package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware is a placeholder that should be replaced with actual middleware
func AuthMiddleware(c *fiber.Ctx) error {
	// This is a placeholder - the actual auth middleware is in the middleware package
	// For now, we'll allow all requests to pass through
	return c.Next()
}