package models

import (
	"time"
)

// Connector represents a connector definition
type Connector struct {
	ID            string    `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	Description   string    `json:"description" db:"description"`
	Type          string    `json:"type" db:"type"`           // rest, soap, db, file, mqtt, etc.
	TenantID      string    `json:"tenant_id" db:"tenant_id"`
	CreatedBy     string    `json:"created_by" db:"created_by"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	Schema        string    `json:"schema" db:"schema"`         // JSON schema for connector configuration
	Actions       []Action  `json:"actions" db:"actions"`       // Available actions for this connector
	Status        string    `json:"status" db:"status"`         // active, inactive, draft
	Version       int       `json:"version" db:"version"`
}

// Action represents an action that a connector can perform
type Action struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Method      string                 `json:"method"`      // get, post, put, delete, etc. for REST
	Endpoint    string                 `json:"endpoint"`    // Endpoint for the action
	InputSchema map[string]interface{} `json:"input_schema"` // JSON schema for input parameters
	OutputSchema map[string]interface{} `json:"output_schema"` // JSON schema for output
	Timeout     int                    `json:"timeout"`     // Timeout in seconds
}

// ConnectorConfig represents the configuration for a connector instance
type ConnectorConfig struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	ConnectorID string                `json:"connector_id"`
	TenantID   string                 `json:"tenant_id"`
	Config     map[string]interface{} `json:"config"`      // Connector-specific configuration
	Status     string                 `json:"status"`      // active, inactive, draft
	CreatedAt  time.Time              `json:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
}

// ConnectorExecution represents the execution of a connector action
type ConnectorExecution struct {
	ID           string    `json:"id"`
	ConnectorID  string    `json:"connector_id"`
	ActionID     string    `json:"action_id"`
	TenantID     string    `json:"tenant_id"`
	Status       string    `json:"status"`       // running, completed, failed
	StartedAt    time.Time `json:"started_at"`
	FinishedAt   time.Time `json:"finished_at"`
	InputData    string    `json:"input_data"`   // JSON input data
	OutputData   string    `json:"output_data"`  // JSON output data
	Error        string    `json:"error"`
	ExecutionTime int64    `json:"execution_time"` // in milliseconds
}

// CreateConnectorRequest represents the request to create a connector
type CreateConnectorRequest struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Type        string                 `json:"type"`
	Schema      map[string]interface{} `json:"schema"`
	Actions     []Action               `json:"actions"`
}

// TestConnectionRequest represents the request to test a connector
type TestConnectionRequest struct {
	Config map[string]interface{} `json:"config"`
}