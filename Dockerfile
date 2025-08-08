# 1. Build stage
FROM golang:1.24.6 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app .

# 2. Final stage
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/app .

# Optional: install sqlite3 CLI for debugging
RUN apt-get update && apt-get install -y --no-install-recommends sqlite3 \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

CMD ["./app"]
