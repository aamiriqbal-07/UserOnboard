# Stage 1: Build the Go binary
FROM golang:1.23.5 as builder

WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code and build the binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Create a minimal image with the Go binary
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Ensure the binary has execution permissions
RUN chmod +x ./main

# Expose the application port
EXPOSE 8080

# Use entrypoint to pass command-line arguments to the binary
ENTRYPOINT ["./main"]
