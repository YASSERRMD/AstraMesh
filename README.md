# Integration + Agentic + GenAI Platform

Next-generation Integration Platform similar to IBM webMethods / MuleSoft but modern and AI-native.

## Features
- Multi-protocol integration (REST, SOAP, gRPC, GraphQL, MQTT, WebSockets, etc.)
- Flow designer with visual editor
- High-throughput data flow execution
- GenAI integration (LLM, embeddings, vision, RAG)
- MCP tool exposure
- Connector framework
- Enterprise HA, scaling, observability
- SvelteKit web control plane

## Tech Stack
- **Backend**: Go 1.22+
- **Frontend**: SvelteKit
- **Database**: PostgreSQL
- **Message Queue**: Kafka/NATS
- **Vector DB**: Qdrant/pgvector
- **Observability**: OpenTelemetry, Loki, Jaeger, Prometheus

## Architecture
The platform consists of Control Plane and Data Plane services for managing and executing integration flows.

## Docker Compose Deployment

This repository includes a complete Docker Compose setup for easy deployment of the entire platform.

### Quick Start

1. Clone the repository:
```bash
git clone <repository-url>
cd <repository-name>
```

2. Set up environment variables:
```bash
cp .env.example .env
# Edit .env to add your API keys
```

3. Start the platform:
```bash
docker-compose up -d
```

4. Check the status:
```bash
docker-compose ps
```

### Services

The Docker Compose setup includes:

- **Backend Services**:
  - `gateway-api`: Main API gateway and control plane
  - `flow-runtime`: Flow execution engine
  - `genai-orchestrator`: AI model orchestration
  - `rag-service`: Retrieval Augmented Generation service
  - `connector-runtime`: Connector execution runtime

- **Infrastructure Services**:
  - `postgres`: PostgreSQL database with vector extension
  - `redis`: Caching and session storage
  - `nats`: Message queuing with JetStream
  - `qdrant`: Vector database for embeddings
  - `minio`: S3-compatible object storage
  - `prometheus`: Metrics collection
  - `grafana`: Visualization dashboard
  - `jaeger`: Distributed tracing
  - `loki`: Log aggregation

- **Frontend**:
  - `frontend`: SvelteKit web application

### Ports

| Service | Port |
|---------|------|
| Frontend | 3001 |
| Gateway API | 8080 |
| PostgreSQL | 5432 |
| Redis | 6379 |
| NATS | 4222 |
| Qdrant | 6333 |
| MinIO Console | 9001 |
| Grafana | 3000 |
| Jaeger UI | 16686 |
| Prometheus | 9090 |
| Loki | 3100 |

### Environment Variables

Create a `.env` file in the root directory with the following variables:

```bash
# API Keys for GenAI services
OPENAI_API_KEY=your_openai_api_key
ANTHROPIC_API_KEY=your_anthropic_api_key
GEMINI_API_KEY=your_gemini_api_key

# Security
JWT_SECRET=your_jwt_secret_key
QDRANT_API_KEY=your_qdrant_api_key

# Optional: Local LLM (Ollama/vLLM)
OLLAMA_BASE_URL=http://host.docker.internal:11434
```

### Scaling

The platform is designed to scale horizontally:

```bash
# Scale flow runtime to handle more concurrent flows
docker-compose up -d --scale flow-runtime=5
```
