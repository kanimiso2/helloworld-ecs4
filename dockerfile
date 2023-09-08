FROM ubuntu:18.04
FROM golang:1.20.5

WORKDIR /app

COPY ./src /app
# Goモジュールをコピーして依存関係を解決
COPY go.mod .
COPY go.sum .

# モジュールをダウンロード
RUN go mod download
EXPOSE 80
CMD ["go", "run", "main.go"]