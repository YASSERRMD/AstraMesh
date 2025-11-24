package models

import (
	"time"
)

// GenAIModel represents a configuration for a GenAI model
type GenAIModel struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Provider    string    `json:"provider" db:"provider"`       // openai, gemini, anthropic, local
	Model       string    `json:"model" db:"model"`            // gpt-4, gpt-3.5-turbo, etc.
	TenantID    string    `json:"tenant_id" db:"tenant_id"`
	APIKey      string    `json:"api_key" db:"api_key"`        // Encrypted in DB
	BaseURL     string    `json:"base_url" db:"base_url"`      // For local models
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	MaxTokens   int       `json:"max_tokens" db:"max_tokens"`
	Temperature float64   `json:"temperature" db:"temperature"`
	Status      string    `json:"status" db:"status"`          // active, inactive
}

// ChatRequest represents a chat completion request
type ChatRequest struct {
	Model     string         `json:"model"`
	Messages  []ChatMessage  `json:"messages"`
	Stream    bool           `json:"stream"`
	MaxTokens int            `json:"max_tokens,omitempty"`
	Temperature float64      `json:"temperature,omitempty"`
	TopP      float64        `json:"top_p,omitempty"`
	Stop      []string       `json:"stop,omitempty"`
	PresencePenalty float64  `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`
}

// ChatMessage represents a message in a chat conversation
type ChatMessage struct {
	Role    string `json:"role"`    // system, user, assistant
	Content string `json:"content"`
}

// ChatResponse represents a chat completion response
type ChatResponse struct {
	ID                string        `json:"id"`
	Object            string        `json:"object"`
	Created           int64         `json:"created"`
	Model             string        `json:"model"`
	Choices           []ChatChoice  `json:"choices"`
	Usage             Usage         `json:"usage"`
	SystemFingerprint string        `json:"system_fingerprint"`
}

// ChatChoice represents a choice in a chat response
type ChatChoice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
}

// Usage represents token usage information
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// EmbeddingRequest represents an embedding request
type EmbeddingRequest struct {
	Model string      `json:"model"`
	Input interface{} `json:"input"`  // string or array of strings
}

// EmbeddingResponse represents an embedding response
type EmbeddingResponse struct {
	Object string        `json:"object"`
	Data   []Embedding   `json:"data"`
	Model  string        `json:"model"`
	Usage  Usage         `json:"usage"`
}

// Embedding represents a single embedding
type Embedding struct {
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
	Index     int       `json:"index"`
}

// VisionRequest represents a vision request
type VisionRequest struct {
	Model  string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
	MaxTokens int         `json:"max_tokens,omitempty"`
}

// PromptTemplate represents a reusable prompt template
type PromptTemplate struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	TenantID    string    `json:"tenant_id" db:"tenant_id"`
	Category    string    `json:"category" db:"category"`      // classification, summarization, etc.
	Template    string    `json:"template" db:"template"`      // Go template syntax
	Variables   []string  `json:"variables" db:"variables"`    // Template variables
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// GenAIExecution represents the execution of a GenAI request
type GenAIExecution struct {
	ID          string    `json:"id"`
	RequestID   string    `json:"request_id"`
	ModelID     string    `json:"model_id"`
	TenantID    string    `json:"tenant_id"`
	RequestType string    `json:"request_type"`  // chat, embed, vision, etc.
	Status      string    `json:"status"`        // running, completed, failed
	StartedAt   time.Time `json:"started_at"`
	FinishedAt  time.Time `json:"finished_at"`
	RequestData string    `json:"request_data"`  // JSON request data
	ResponseData string   `json:"response_data"` // JSON response data
	Error       string    `json:"error"`
	ExecutionTime int64   `json:"execution_time"` // in milliseconds
	Usage       Usage     `json:"usage"`
}