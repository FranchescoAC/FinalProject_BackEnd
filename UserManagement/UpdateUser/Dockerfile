# Stage 1: Build
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application in release mode
RUN go build -o main .

# Stage 2: Run
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy compiled binary from the builder stage
COPY --from=builder /app/main .

# Set environment variable for the port
ENV PORT=2003

# Expose the service port
EXPOSE ${PORT}

# Run the application
CMD ["/app/main"]
