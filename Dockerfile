# -----------------------------
# Stage 1: Build the Go binary
# -----------------------------
FROM golang:1.24 AS builder

WORKDIR /app

# Copy the Go code
COPY main.go .

# Build a statically linked binary for Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# -----------------------------
# Stage 2: Minimal runtime image
# -----------------------------
FROM alpine:latest

WORKDIR /root/

# Add certificates (optional, if HTTPS is needed)
RUN apk add --no-cache ca-certificates

# Copy the binary
COPY --from=builder /app/main .

# Expose port
EXPOSE 8081

# Run the app
CMD ["./main"]
