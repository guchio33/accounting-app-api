# Build Stage
FROM golang:alpine3.20 AS builder

WORKDIR /app

# go.mod と go.sum をコピーして依存関係をダウンロード
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# 必要なツールをインストール
RUN go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/air-verse/air@latest

# ソースコードをコピー
COPY . .

CMD ["go", "run", "./cmd/server/main.go"]