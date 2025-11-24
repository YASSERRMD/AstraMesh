package models

import (
	"time"
)

// Document represents a document in the RAG system
type Document struct {
	ID          string            `json:"id" db:"id"`
	Title       string            `json:"title" db:"title"`
	Content     string            `json:"content" db:"content"`
	TenantID    string            `json:"tenant_id" db:"tenant_id"`
	Source      string            `json:"source" db:"source"`         // file path, URL, etc.
	MimeType    string            `json:"mime_type" db:"mime_type"`   // text/plain, application/pdf, etc.
	Metadata    map[string]interface{} `json:"metadata" db:"metadata"` // Additional metadata
	CreatedAt   time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at" db:"updated_at"`
	Status      string            `json:"status" db:"status"`         // processing, processed, failed
	ChunkCount  int               `json:"chunk_count" db:"chunk_count"`
}

// DocumentChunk represents a chunk of a document
type DocumentChunk struct {
	ID          string    `json:"id" db:"id"`
	DocumentID  string    `json:"document_id" db:"document_id"`
	Content     string    `json:"content" db:"content"`
	ChunkIndex  int       `json:"chunk_index" db:"chunk_index"`
	EmbeddingID string    `json:"embedding_id" db:"embedding_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// EmbeddingVector represents an embedding vector stored in the vector DB
type EmbeddingVector struct {
	ID          string    `json:"id" db:"id"`
	Vector      []float64 `json:"vector" db:"vector"`  // The actual embedding vector
	DocumentID  string    `json:"document_id" db:"document_id"`
	ChunkID     string    `json:"chunk_id" db:"chunk_id"`
	Model       string    `json:"model" db:"model"`    // The model used to generate the embedding
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	Metadata    map[string]interface{} `json:"metadata" db:"metadata"` // Additional metadata for filtering
}

// RAGSearchRequest represents a search request in the RAG system
type RAGSearchRequest struct {
	Query       string                 `json:"query"`
	TopK        int                    `json:"top_k"`           // Number of results to return
	Filters     map[string]interface{} `json:"filters"`         // Metadata filters
	MinScore    float64                `json:"min_score"`       // Minimum similarity score
	Model       string                 `json:"model"`           // Model to use for query embedding
	MaxTokens   int                    `json:"max_tokens"`      // Maximum tokens in response
	Temperature float64                `json:"temperature"`     // Temperature for response generation
}

// RAGSearchResult represents a search result from the RAG system
type RAGSearchResult struct {
	DocumentID  string    `json:"document_id"`
	ChunkID     string    `json:"chunk_id"`
	Content     string    `json:"content"`
	Similarity  float64   `json:"similarity"`
	Metadata    map[string]interface{} `json:"metadata"`
	Source      string    `json:"source"`
	CreatedAt   time.Time `json:"created_at"`
}

// RAGSearchResponse represents the response from a RAG search
type RAGSearchResponse struct {
	Query        string            `json:"query"`
	Results      []RAGSearchResult `json:"results"`
	Response     string            `json:"response"`          // Generated response based on context
	ExecutionTime int64            `json:"execution_time"`    // in milliseconds
}

// RAGExecution represents the execution of a RAG request
type RAGExecution struct {
	ID          string    `json:"id"`
	RequestID   string    `json:"request_id"`
	TenantID    string    `json:"tenant_id"`
	Status      string    `json:"status"`        // running, completed, failed
	StartedAt   time.Time `json:"started_at"`
	FinishedAt  time.Time `json:"finished_at"`
	Query       string    `json:"query"`         // The original query
	Context     string    `json:"context"`       // Retrieved context
	Response    string    `json:"response"`      // Generated response
	Error       string    `json:"error"`
	ExecutionTime int64   `json:"execution_time"` // in milliseconds
	SearchResults []RAGSearchResult `json:"search_results"`
}