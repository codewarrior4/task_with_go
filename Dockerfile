# Use the official Golang image as the base
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from a smaller base image
FROM alpine:latest  

# Install necessary dependencies (if any, e.g., ca-certificates)
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the previous stage
COPY --from=builder /app/main .

# Expose port the app runs on
EXPOSE 4023

# Command to run the executable
CMD ["./main"]
