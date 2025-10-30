# ---- Build stage ----
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compile the application for Linux
# CGO_ENABLED=0 is important for static builds
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# ---- Final stage ----
# Use a minimal base image
FROM alpine:latest

WORKDIR /root/

# Copy the compiled binary from the build stage
COPY --from=builder /app/main .

# Copy email templates and other assets if needed at runtime
COPY --from=builder /app/templates ./templates

EXPOSE 8080

CMD ["./main"]