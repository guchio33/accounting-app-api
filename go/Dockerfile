# Build Stage
FROM golang:alpine3.20 AS builder

WORKDIR /app

# go.mod と go.sum をコピーして依存関係をダウンロード
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# 必要なツールをインストール (開発時のみ)
RUN go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/air-verse/air@latest

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o /app/tmp/main ./cmd/server/main.go

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]