
# 使用 Golang 镜像作为构建阶段
FROM golang:1.23.0 AS builder

# 设置目录
RUN mkdir -p /app /app/health

# 设置工作目录
WORKDIR /app
COPY . .

RUN env CGO_ENABLED="0" GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o /app/server ./cmd/server

# 最终阶段
FROM gcr.io/distroless/static-debian11

ENV GIN_MODE=release
ENV GOPROXY=https://goproxy.cn,direct

COPY --from=builder /app/server /

EXPOSE 9999

CMD ["./server"]
