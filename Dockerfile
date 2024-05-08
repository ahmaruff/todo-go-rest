# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Ahmad <ahmadmaruf2701@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Disable Go Modules
ENV GO111MODULE=off

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 4000

# Command to run the executable
CMD ["./main"]