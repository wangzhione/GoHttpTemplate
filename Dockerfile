# 使用 Alpine 作为构建基础
FROM golang:alpine AS builder

LABEL stage=gobuilder

# 关闭 CGO，减少依赖
ENV CGO_ENABLED=0

# 二进制文件名，可通过 --build-arg BINARY_NAME=xxx 指定
ARG BINARY_NAME=gohttptemplate

# 安装必要的系统工具
RUN apk update --no-cache && apk add --no-cache tzdata ca-certificates

WORKDIR /build

# 复制依赖文件并缓存 go mod 下载，提高构建速度
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

# 复制所有源码
COPY . .

# 编译 Go 可执行文件
# go tool addr2line -e /app/${BINARY_NAME} {ptr}
RUN go build -ldflags="-s -w" -o /app/${BINARY_NAME} main.go

# 使用最小 scratch 镜像
FROM scratch

# 复制证书和时区文件
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/

# 设置时区
ENV TZ=Asia/Shanghai

WORKDIR /app

# 复制可执行文件并赋予执行权限
COPY --from=builder /app/${BINARY_NAME} /app/${BINARY_NAME}

# 复制配置文件
COPY --from=builder /build/resource/etc/prod.toml /app/resource/etc/prod.toml

# 声明开放端口 8089, 硬编码, 自行手工修改
EXPOSE 8089

ENTRYPOINT ["/app/${BINARY_NAME}"]

# 允许动态 "-f" 参数传入日志路径，默认走程序 internal/inits/config.go 内部写死的 resource/etc/prod.toml
