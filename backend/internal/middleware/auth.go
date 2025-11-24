package middleware

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("default-secret-key-change-in-production") // This should come from config

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "Authorization header required",
		})
	}

	tokenString := ""
	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString = strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid authorization header format",
		})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}

	// Add claims to context for use in handlers
	c.Locals("user_id", claims["user_id"])
	c.Locals("tenant_id", claims["tenant_id"])
	c.Locals("role", claims["role"])

	return c.Next()
}

func APIKeyMiddleware(c *fiber.Ctx) error {
	apiKey := c.Get("X-API-Key")
	if apiKey == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "X-API-Key header required",
		})
	}

	// In a real implementation, you would verify the API key against a database
	// For now, we'll just check if it's present
	if len(apiKey) < 10 { // Basic validation
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid API key",
		})
	}

	return c.Next()
}