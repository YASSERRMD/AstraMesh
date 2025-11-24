-- Database schema for the Integration + Agentic + GenAI Platform

-- Tenants table
CREATE TABLE tenants (
    id VARCHAR(64) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Users table
CREATE TABLE users (
    id VARCHAR(64) PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    tenant_id VARCHAR(64) NOT NULL REFERENCES tenants(id),
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    role VARCHAR(50) DEFAULT 'user',
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Roles table (for RBAC)
CREATE TABLE roles (
    id VARCHAR(64) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    permissions JSONB, -- JSON array of permissions
    tenant_id VARCHAR(64) NOT NULL REFERENCES tenants(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Flows table
CREATE TABLE flows (
    id VARCHAR(64) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    tenant_id VARCHAR(64) NOT NULL REFERENCES tenants(id),
    created_by VARCHAR(64) NOT NULL REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    flow_data JSONB, -- JSON representation of the flow
    status VARCHAR(20) DEFAULT 'draft', -- draft, active, inactive
    version INTEGER DEFAULT 1
);

-- Flow versions table (for versioning)
CREATE TABLE flow_versions (
    id VARCHAR(64) PRIMARY KEY,
    flow_id VARCHAR(64) NOT NULL REFERENCES flows(id),
    version INTEGER NOT NULL,
    flow_data JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(64) NOT NULL REFERENCES users(id)
);

-- Connectors table
CREATE TABLE connectors (
    id VARCHAR(64) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    type VARCHAR(50) NOT NULL, -- rest, soap, db, file, mqtt, etc.
    tenant_id VARCHAR(64) NOT NULL REFERENCES tenants(id),
    created_by VARCHAR(64) NOT NULL REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    schema JSONB, -- JSON schema for connector configuration
    actions JSONB, -- Array of available actions
    status VARCHAR(20) DEFAULT 'draft', -- draft, active, inactive
    version INTEGER DEFAULT 1
);

-- Connector types table (for predefined connector templates)
CREATE TABLE connector_types (
    id VARCHAR(64) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    type VARCHAR(50) NOT NULL, -- rest, soap, db, file, mqtt, etc.
    schema JSONB, -- JSON schema for connector configuration
    actions JSONB, -- Array of available actions
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Executions table (for flow executions)
CREATE TABLE executions (
    id VARCHAR(64) PRIMARY KEY,
    flow_id VARCHAR(64) NOT NULL REFERENCES flows(id),
    tenant_id VARCHAR(64) NOT NULL REFERENCES tenants(id),
    status VARCHAR(20) DEFAULT 'running', -- running, completed, failed, cancelled
    started_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    finished_at TIMESTAMP,
    input_data JSONB,
    output_data JSONB,
    error TEXT,
    trace_id VARCHAR(64),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- GenAI models table
CREATE TABLE genai_models (
    id VARCHAR(64) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    provider VARCHAR(50) NOT NULL, -- openai, gemini, anthropic, local
    model VARCHAR(100) NOT NULL, -- gpt-4, gpt-3.5-turbo, etc.
    tenant_id VARCHAR(64) NOT NULL REFERENCES tenants(id),
    api_key_encrypted TEXT, -- Encrypted API key
    base_url VARCHAR(500), -- For local models
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    max_tokens INTEGER DEFAULT 4096,
    temperature DECIMAL(3,2) DEFAULT 0.7,
    status VARCHAR(20) DEFAULT 'active' -- active, inactive
);

-- Prompt templates table
CREATE TABLE prompt_templates (
    id VARCHAR(64) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    tenant_id VARCHAR(64) NOT NULL REFERENCES tenants(id),
    category VARCHAR(100), -- classification, summarization, etc.
    template TEXT NOT NULL, -- Go template syntax
    variables JSONB, -- Array of template variables
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Documents table (for RAG)
CREATE TABLE documents (
    id VARCHAR(64) PRIMARY KEY,
    title VARCHAR(500) NOT NULL,
    content TEXT,
    tenant_id VARCHAR(64) NOT NULL REFERENCES tenants(id),
    source VARCHAR(500), -- file path, URL, etc.
    mime_type VARCHAR(100), -- text/plain, application/pdf, etc.
    metadata JSONB, -- Additional metadata
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(20) DEFAULT 'processing', -- processing, processed, failed
    chunk_count INTEGER DEFAULT 0
);

-- Embeddings metadata table (for RAG)
CREATE TABLE embeddings_metadata (
    id VARCHAR(64) PRIMARY KEY,
    document_id VARCHAR(64) NOT NULL REFERENCES documents(id),
    chunk_id VARCHAR(64),
    model VARCHAR(100) NOT NULL, -- The model used to generate the embedding
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    metadata JSONB, -- Additional metadata for filtering
    tenant_id VARCHAR(64) NOT NULL REFERENCES tenants(id)
);

-- MCP tools metadata table
CREATE TABLE mcp_tools_metadata (
    id VARCHAR(64) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    flow_id VARCHAR(64) NOT NULL REFERENCES flows(id),
    tenant_id VARCHAR(64) NOT NULL REFERENCES tenants(id),
    schema JSONB, -- JSON schema for the tool
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(20) DEFAULT 'active', -- active, inactive
    target_ids JSONB -- Array of target IDs this tool is exposed to
);

-- Indexes for performance
CREATE INDEX idx_flows_tenant_id ON flows(tenant_id);
CREATE INDEX idx_flows_status ON flows(status);
CREATE INDEX idx_executions_flow_id ON executions(flow_id);
CREATE INDEX idx_executions_tenant_id ON executions(tenant_id);
CREATE INDEX idx_executions_status ON executions(status);
CREATE INDEX idx_executions_started_at ON executions(started_at);
CREATE INDEX idx_documents_tenant_id ON documents(tenant_id);
CREATE INDEX idx_documents_status ON documents(status);
CREATE INDEX idx_mcp_tools_tenant_id ON mcp_tools_metadata(tenant_id);
CREATE INDEX idx_users_tenant_id ON users(tenant_id);
CREATE INDEX idx_connectors_tenant_id ON connectors(tenant_id);
CREATE INDEX idx_genai_models_tenant_id ON genai_models(tenant_id);