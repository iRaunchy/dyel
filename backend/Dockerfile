# backend/Dockerfile

# Builder stage
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Only copy the module files from backend/
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the backend source
COPY . .

# Build the binary
RUN go build -o dyel main.go

# Final stage
FROM alpine:latest
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/dyel /usr/local/bin/dyel
ENTRYPOINT ["dyel"]