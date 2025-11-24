#!/bin/bash

# Integration Platform Startup Script

set -e  # Exit on any error

echo "🚀 Starting Integration Platform..."

# Check if Docker is available
if ! command -v docker &> /dev/null; then
    echo "❌ Docker is not installed or not in PATH"
    exit 1
fi

# Check if docker-compose is available
if ! command -v docker-compose &> /dev/null; then
    # Try docker compose (v2)
    if ! command -v docker compose &> /dev/null; then
        echo "❌ Docker Compose is not installed or not in PATH"
        exit 1
    fi
fi

# Check if .env file exists, if not create from example
if [ ! -f .env ]; then
    if [ -f .env.example ]; then
        echo "📝 Copying .env.example to .env"
        cp .env.example .env
        echo "⚠️  Please edit .env to set your API keys and configuration"
    else
        echo "⚠️  No .env file found. Please create one before running."
        exit 1
    fi
fi

echo "🔧 Starting services..."

# Use the appropriate docker-compose command
if command -v docker-compose &> /dev/null; then
    docker-compose up -d
else
    docker compose up -d
fi

echo "⏳ Waiting for services to start..."

# Wait for key services to be healthy
echo "🏥 Checking service health..."

# Wait up to 60 seconds for services to be healthy
timeout 60 bash -c '
  while [[ $(docker-compose ps -q | xargs -I {} docker inspect {} --format="{{.State.Health.Status}}") =~ "starting" ]]; do
    echo "⏳ Waiting for services to be healthy..."
    sleep 5
  done
'

echo "✅ Integration Platform is now running!"

echo ""
echo "📋 Services are available at:"
echo "   Frontend: http://localhost:3001"
echo "   Gateway API: http://localhost:8080"
echo "   Grafana: http://localhost:3000 (admin/admin)"
echo "   Prometheus: http://localhost:9090"
echo "   Jaeger: http://localhost:16686"
echo "   MinIO Console: http://localhost:9001 (minioadmin/minioadmin123)"
echo ""

echo "📊 To view logs in real-time:"
echo "   docker-compose logs -f"
echo ""

echo "🔄 To stop the platform:"
echo "   docker-compose down"
echo ""

echo "🔧 For development mode with hot-reload:"
echo "   docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d"
echo ""