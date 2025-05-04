# ── Stage 1: Build the Go binary ────────────────────────────────────────────
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy code & build
COPY . .
RUN go build -o dyel main.go

# ── Stage 2: Create minimal runtime image ──────────────────────────────────
FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /root/
# Copy binary from builder
COPY --from=builder /app/dyel .
# Copy .env if you prefer baking it in (optional, you can also mount it)
COPY .env .

EXPOSE 8080
CMD ["./dyel"]
