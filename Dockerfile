# Dockerfile
# 构建阶段
FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o gin-api-template .

# 运行阶段
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /build/gin-api-template .
COPY .env .

EXPOSE 8080

CMD ["./gin-api-template"]