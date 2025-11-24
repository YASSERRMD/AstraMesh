package models

import (
	"time"
)

// MCPTarget represents an MCP target (like VS Code, Cursor, etc.)
type MCPTarget struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	TenantID    string    `json:"tenant_id" db:"tenant_id"`
	Type        string    `json:"type" db:"type"`              // stdio, websocket
	Endpoint    string    `json:"endpoint" db:"endpoint"`      // For websocket connections
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Status      string    `json:"status" db:"status"`          // connected, disconnected, error
}

// MCPTool represents an MCP tool (exposed flow)
type MCPTool struct {
	ID          string                 `json:"id" db:"id"`
	Name        string                 `json:"name" db:"name"`
	Description string                 `json:"description" db:"description"`
	FlowID      string                 `json:"flow_id" db:"flow_id"`
	TenantID    string                 `json:"tenant_id" db:"tenant_id"`
	Schema      map[string]interface{} `json:"schema" db:"schema"`       // JSON schema for the tool
	CreatedAt   time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at" db:"updated_at"`
	Status      string                 `json:"status" db:"status"`       // active, inactive
	TargetIDs   []string               `json:"target_ids" db:"target_ids"` // Which targets this tool is exposed to
}

// MCPToolCall represents a call to an MCP tool
type MCPToolCall struct {
	ID          string    `json:"id"`
	ToolID      string    `json:"tool_id"`
	TargetID    string    `json:"target_id"`
	TenantID    string    `json:"tenant_id"`
	CallID      string    `json:"call_id"`       // MCP call ID
	Arguments   string    `json:"arguments"`     // JSON arguments for the tool call
	Status      string    `json:"status"`        // pending, running, completed, failed
	StartedAt   time.Time `json:"started_at"`
	FinishedAt  time.Time `json:"finished_at"`
	Result      string    `json:"result"`        // JSON result of the tool call
	Error       string    `json:"error"`
	ExecutionTime int64   `json:"execution_time"` // in milliseconds
}

// MCPConfig represents the configuration for MCP integration
type MCPConfig struct {
	ID          string    `json:"id" db:"id"`
	TenantID    string    `json:"tenant_id" db:"tenant_id"`
	ServerName  string    `json:"server_name" db:"server_name"`  // Name of the MCP server
	Description string    `json:"description" db:"description"`  // Description of the MCP server
	Endpoint    string    `json:"endpoint" db:"endpoint"`        // Endpoint for the MCP server
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Status      string    `json:"status" db:"status"`            // active, inactive
	Tools       []string  `json:"tools" db:"tools"`              // IDs of tools exposed by this server
}

// MCPToolSchema represents the schema for an MCP tool
type MCPToolSchema struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	InputSchema map[string]interface{} `json:"input_schema"`  // JSON schema for input parameters
	OutputSchema map[string]interface{} `json:"output_schema"` // JSON schema for output
	Examples    []map[string]interface{} `json:"examples"`    // Example inputs/outputs
}

// MCPToolRequest represents a request to call an MCP tool
type MCPToolRequest struct {
	ToolID    string                 `json:"tool_id"`
	Arguments map[string]interface{} `json:"arguments"`
	TargetID  string                 `json:"target_id"`
}

// MCPToolResponse represents a response from an MCP tool
type MCPToolResponse struct {
	ID     string      `json:"id"`
	Result interface{} `json:"result"`
	Error  string      `json:"error,omitempty"`
}