# Use the official golang image as the base
FROM golang:1.22 AS builder

# Set working directory for the build stage
WORKDIR /app

# Copy your Go source code and dependencies (go.mod, go.sum)
COPY . .

# Download Go dependencies
RUN go mod download

# Build the Go binary (assuming main.go is your entrypoint)
RUN go build -o todo .

# Use a smaller image for running the application
FROM alpine:latest AS runner

# Copy the application binary from the build stage
COPY --from=builder /app/todo /app/todo

# Set working directory for the runner stage
WORKDIR /app

# Expose the port your application listens on (replace 8080 with your actual port)
EXPOSE 8443

# Run your application binary
CMD ["./todo"]
