# -----------------------------
# Build stage
# -----------------------------
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source files
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# -----------------------------
# Final stage
# -----------------------------
FROM alpine:latest

WORKDIR /root/

# Install CA certificates (optional)
RUN apk add --no-cache ca-certificates

# Copy built binary
COPY --from=builder /app/main .

# Expose port
EXPOSE 8081

# Run app
CMD ["./main"]