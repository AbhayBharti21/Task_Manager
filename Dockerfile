# Use official Go image
FROM golang:1.25-alpine

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all source code
COPY . .

# Build the application
RUN go build -o main ./cmd/main.go

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./main"]
