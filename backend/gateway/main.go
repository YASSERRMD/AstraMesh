package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"platform-backend/internal/config"
	"platform-backend/internal/routes"
	"platform-backend/internal/middleware"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize configuration
	cfg := config.LoadConfig()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(middleware.RequestIDMiddleware)
	app.Use(middleware.RateLimiterMiddleware(cfg.RateLimit))

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"service": "gateway",
		})
	})

	// Register routes
	routes.SetupRoutes(app)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Gateway service starting on port %s", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}