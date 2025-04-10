# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /task

# Install required packages (bash + ca-certificates for HTTPS)
RUN apk --no-cache add ca-certificates bash netcat-openbsd

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Minimal runtime image
FROM alpine:latest

# Install required packages (bash + ca-certificates for HTTPS)
RUN apk --no-cache add ca-certificates bash

# Set working directory
WORKDIR /root/

# Copy binary and wait script
COPY --from=builder /task/main .
COPY wait-for-db.sh /wait-for-db.sh
RUN chmod +x /wait-for-db.sh

# Expose the app port
EXPOSE 4023

# Run the app after ensuring DB is up
CMD ["/wait-for-db.sh", "db:3306", "--", "./main"]
