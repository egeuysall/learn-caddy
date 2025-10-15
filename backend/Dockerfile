# Build stage
FROM golang:1.25.2-alpine AS builder

# Install git (needed for go modules)
RUN apk add --no-cache git

WORKDIR /app

# Download dependencies separately for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code and build
COPY . .
ARG BUILD_VERSION=1.0.0
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -ldflags="-s -w -X 'main.version=${BUILD_VERSION}'" -o api ./main.go

# Runtime stage
FROM alpine:3.20

# Install certificates for HTTPS requests
RUN apk add --no-cache ca-certificates

# Add a non-root user
RUN adduser -D -g '' appuser
USER appuser

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/api .

# Run the binary
CMD ["./api"]
