FROM golang:1.19-alpine as builder

# 设置工作目录
WORKDIR /app

# 复制源代码
COPY . .

# 编译 Go 程序
RUN go mod vendor \
  && go build -o tmp_cleaner main.go

# 使用一个更小的镜像来运行程序
FROM alpine:latest as runner

# 安装 dumb-init
RUN apk add --no-cache dumb-init

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/tmp_cleaner /app/tmp_cleaner

# 使用 dumb-init 作为入口点
ENTRYPOINT ["dumb-init", "--"]

ENV TMP_DIR=/tmp \
 DAYS= \
 INTERVAL=

# 运行程序
CMD ["/app/tmp_cleaner"]