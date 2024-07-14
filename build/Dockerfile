# Backend Dockerfile
FROM golang:1.22-alpine

WORKDIR /app

# Install OpenSSL for generating certificates
RUN apk add --no-cache openssl

# Copy and download Go dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the application code
COPY . .

# Generate SSL certificates
RUN openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes -subj "/CN=localhost"

# Build the Go application
RUN go build -o main .

# Expose the port
EXPOSE 443

# Command to run the application
CMD ["./main"]
