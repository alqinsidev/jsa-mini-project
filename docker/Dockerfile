# Stage 1: Build the Go application
FROM golang:1.18-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests and download dependencies
COPY src/go.mod src/go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY src/ ./

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./app

# Stage 2: Create the final image
FROM alpine:3.14

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main ./

# Command to run the executable
CMD ["./main"]
