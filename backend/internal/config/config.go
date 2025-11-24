package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	ServerPort    string
	DatabaseURL   string
	RedisURL      string
	KafkaBrokers  []string
	NATSServer    string
	JaegerAgent   string
	PrometheusURL string
	RateLimit     int
	AuthSecret    string
	OpenAIKey     string
	GeminiKey     string
	AnthropicKey  string
	QdrantURL     string
}

func LoadConfig() *Config {
	return &Config{
		ServerPort:    getEnv("SERVER_PORT", "8080"),
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/platform?sslmode=disable"),
		RedisURL:      getEnv("REDIS_URL", "redis://localhost:6379"),
		KafkaBrokers:  []string{getEnv("KAFKA_BROKERS", "localhost:9092")},
		NATSServer:    getEnv("NATS_SERVER", "nats://localhost:4222"),
		JaegerAgent:   getEnv("JAEGER_AGENT", "localhost:6831"),
		PrometheusURL: getEnv("PROMETHEUS_URL", "http://localhost:9090"),
		RateLimit:     getEnvAsInt("RATE_LIMIT", 100),
		AuthSecret:    getEnv("AUTH_SECRET", "default-secret-key-change-in-production"),
		OpenAIKey:     getEnv("OPENAI_KEY", ""),
		GeminiKey:     getEnv("GEMINI_KEY", ""),
		AnthropicKey:  getEnv("ANTHROPIC_KEY", ""),
		QdrantURL:     getEnv("QDRANT_URL", "http://localhost:6334"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		log.Printf("Invalid value for %s, using default: %d", key, defaultValue)
	}
	return defaultValue
}