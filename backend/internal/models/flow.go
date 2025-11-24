package models

import (
	"time"

	"github.com/google/uuid"
)

// Flow represents an integration flow
type Flow struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	TenantID    string    `json:"tenant_id" db:"tenant_id"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	FlowData    string    `json:"flow_data" db:"flow_data"` // JSON representation of the flow
	Status      string    `json:"status" db:"status"`       // active, inactive, draft
	Version     int       `json:"version" db:"version"`
}

// FlowNode represents a node in the flow
type FlowNode struct {
	ID       string                 `json:"id"`
	Type     string                 `json:"type"`     // rest, soap, grpc, graphql, mqtt, db, genai, etc.
	Config   map[string]interface{} `json:"config"`   // Configuration specific to the node type
	Input    map[string]interface{} `json:"input"`    // Input mappings
	Output   map[string]interface{} `json:"output"`   // Output mappings
	Next     []string               `json:"next"`     // IDs of next nodes
	ErrorTo  string                 `json:"error_to"` // ID of error handling node
}

// FlowExecution represents a single execution of a flow
type FlowExecution struct {
	ID          string    `json:"id"`
	FlowID      string    `json:"flow_id"`
	TenantID    string    `json:"tenant_id"`
	Status      string    `json:"status"`      // running, completed, failed, cancelled
	StartedAt   time.Time `json:"started_at"`
	FinishedAt  time.Time `json:"finished_at"`
	InputData   string    `json:"input_data"`  // JSON input data
	OutputData  string    `json:"output_data"` // JSON output data
	Error       string    `json:"error"`
	TraceID     string    `json:"trace_id"`
	NodeResults []NodeResult `json:"node_results"`
}

// NodeResult represents the result of executing a single node
type NodeResult struct {
	NodeID      string    `json:"node_id"`
	Status      string    `json:"status"`      // success, failed, skipped
	StartedAt   time.Time `json:"started_at"`
	FinishedAt  time.Time `json:"finished_at"`
	InputData   string    `json:"input_data"`
	OutputData  string    `json:"output_data"`
	Error       string    `json:"error"`
	ExecutionTime int64   `json:"execution_time"` // in milliseconds
}

// CreateFlowRequest represents the request to create a flow
type CreateFlowRequest struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	FlowData    interface{} `json:"flow_data"` // The actual flow definition
}

// UpdateFlowRequest represents the request to update a flow
type UpdateFlowRequest struct {
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	FlowData    interface{} `json:"flow_data,omitempty"`
	Status      string      `json:"status,omitempty"`
}