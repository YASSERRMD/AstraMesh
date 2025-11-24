package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the system
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	TenantID string `json:"tenant_id"`
	Role     string `json:"role"`
}

// LoginRequest represents the login request body
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse represents the login response body
type LoginResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	User         User      `json:"user"`
}

var jwtSecret = []byte("default-secret-key-change-in-production") // This should come from config
var users = []User{ // In a real app, this would be in a database
	{
		ID:       "1",
		Email:    "admin@example.com",
		Password: hashPassword("password123"), // Hashed password
		TenantID: "tenant1",
		Role:     "admin",
	},
}

func hashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Find user by email
	var user *User
	for _, u := range users {
		if u.Email == req.Email {
			user = &u
			break
		}
	}

	if user == nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Generate JWT tokens
	accessToken, err := generateToken(user.ID, user.TenantID, user.Role, time.Hour)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not generate access token",
		})
	}

	refreshToken, err := generateToken(user.ID, user.TenantID, user.Role, 24*time.Hour)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not generate refresh token",
		})
	}

	resp := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(time.Hour),
		User:         *user,
	}

	return c.JSON(resp)
}

func generateToken(userID, tenantID, role string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"tenant_id": tenantID,
		"role":     role,
		"exp":      time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func Register(c *fiber.Ctx) error {
	return c.Status(501).JSON(fiber.Map{
		"error": "Not implemented",
	})
}

func RefreshToken(c *fiber.Ctx) error {
	return c.Status(501).JSON(fiber.Map{
		"error": "Not implemented",
	})
}

func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
		"time":   time.Now().Unix(),
	})
}