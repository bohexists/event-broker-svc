# Build stage
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o event-broker-svc ./cmd/main.go

# Run stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/event-broker-svc .

ENV NATS_URL=nats://nats:4222

CMD ["./event-broker-svc"]