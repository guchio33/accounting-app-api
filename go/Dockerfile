# Build Stage
FROM golang:alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

CMD ["go", "run", "main.go"]