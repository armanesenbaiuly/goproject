FROM golang:1.24.2 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go

FROM ubuntu:22.04

WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/internal/config/db/migrations /app/internal/config/db/migrations
COPY .env .

EXPOSE 8080
CMD ["./main"]