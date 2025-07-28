# Build stage
FROM golang:1.21-alpine AS builder

# Install required system dependencies for Ebiten
RUN apk add --no-cache \
    gcc \
    musl-dev \
    xvfb \
    libx11-dev \
    libxrandr-dev \
    libxinerama-dev \
    libxcursor-dev \
    libxi-dev

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o flappy-bird .

# Runtime stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache \
    xvfb \
    libx11 \
    libxrandr \
    libxinerama \
    libxcursor \
    libxi \
    ca-certificates

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/flappy-bird .

# Create a non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Change ownership of the app directory
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port (if needed for future web version)
EXPOSE 8080

# Run the game
CMD ["./flappy-bird"] 