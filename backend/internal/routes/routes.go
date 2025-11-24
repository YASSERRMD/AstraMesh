package routes

import (
	"github.com/gofiber/fiber/v2"

	"platform-backend/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	// Public routes
	public := app.Group("/api/v1")
	
	// Authentication routes
	auth := public.Group("/auth")
	auth.Post("/login", handlers.Login)
	auth.Post("/register", handlers.Register)
	auth.Post("/refresh", handlers.RefreshToken)

	// Health check
	public.Get("/health", handlers.HealthCheck)

	// Protected routes
	protected := app.Group("/api/v1", handlers.AuthMiddleware) // Using the auth middleware
	
	// Flow management routes
	flows := protected.Group("/flows")
	flows.Get("/", handlers.GetFlows)
	flows.Get("/:id", handlers.GetFlow)
	flows.Post("/", handlers.CreateFlow)
	flows.Put("/:id", handlers.UpdateFlow)
	flows.Delete("/:id", handlers.DeleteFlow)
	flows.Post("/:id/execute", handlers.ExecuteFlow)

	// Connector management routes
	connectors := protected.Group("/connectors")
	connectors.Get("/", handlers.GetConnectors)
	connectors.Get("/:id", handlers.GetConnector)
	connectors.Post("/", handlers.CreateConnector)
	connectors.Put("/:id", handlers.UpdateConnector)
	connectors.Delete("/:id", handlers.DeleteConnector)
	connectors.Post("/:id/test", handlers.TestConnector)

	// GenAI routes
	genai := protected.Group("/genai")
	genai.Post("/chat", handlers.ChatCompletion)
	genai.Post("/embed", handlers.Embedding)
	genai.Post("/vision", handlers.Vision)
	genai.Post("/moderate", handlers.Moderation)

	// RAG routes
	rag := protected.Group("/rag")
	rag.Post("/upload", handlers.UploadDocument)
	rag.Post("/search", handlers.SearchDocuments)
	rag.Get("/documents", handlers.GetDocuments)

	// MCP routes
	mcp := protected.Group("/mcp")
	mcp.Get("/tools", handlers.GetMCPTools)
	mcp.Post("/expose/:flowId", handlers.ExposeAsMCP)

	// Monitoring routes
	monitoring := protected.Group("/monitoring")
	monitoring.Get("/metrics", handlers.GetMetrics)
	monitoring.Get("/logs", handlers.GetLogs)
	monitoring.Get("/health", handlers.HealthCheckDetailed)
}