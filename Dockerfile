# Step 1: Build the Go app
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN go build -o main .

# Step 2: Run the app in a small image
FROM alpine:latest
WORKDIR /root/

# Copy built binary
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
