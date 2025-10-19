
FROM golang:1.25-alpine

WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod ./go.mod
COPY go.sum ./go.sum

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main ./cmd/main.go

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
