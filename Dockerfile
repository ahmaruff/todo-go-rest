# Use the official Golang 1.21 image to create a build artifact.
FROM golang:1.21 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app. Compile statically linked version of the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o todo-go .

# Use a distroless image for the runtime environment to minimize size and surface area of attack
FROM gcr.io/distroless/static-debian11

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/todo-go .

# Expose the port your app runs on
EXPOSE 4000

# Run the binary
CMD ["./todo-go"]