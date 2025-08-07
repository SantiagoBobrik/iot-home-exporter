# 1. Build stage
FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app .

# 2. Final stage
FROM debian:bookworm-slim

WORKDIR /app

# Copy binary and any other assets needed
COPY --from=builder /app/app .
COPY --from=builder /app/mydb.sqlite ./mydb.sqlite

# Install sqlite3 CLI (optional for debugging)
RUN apt-get update && apt-get install -y sqlite3 && apt-get clean

CMD ["./app"]
