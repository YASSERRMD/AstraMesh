# Multi-stage build for the Go backend services

# Build stage
FROM golang:1.22-alpine AS builder

# Install git and other build dependencies
RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

# Allow building from either the repository root (SRC_DIR=backend) or the
# backend directory itself (default).
ARG SRC_DIR=.

# Copy go mod files
COPY ${SRC_DIR}/go.mod ${SRC_DIR}/go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY ${SRC_DIR}/ ./

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Copy any configuration files if needed
COPY --from=builder /app/config/ ./config/ || true

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1

# Run the binary
CMD ["./main"]