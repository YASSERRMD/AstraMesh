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
