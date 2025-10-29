# ----- Stage 1: Builder -----
# ⭐ 修正：使用你 go.mod 中指定的 1.25 版本
FROM golang:1.25-alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    TZ=Asia/Shanghai

# 安装依赖：
# tzdata: 用于数据库连接的 "loc=Asia/Shanghai" 时区支持
# ca-certificates: 用于 go mod download 从 "https://goproxy.cn" (HTTPS) 下载
RUN apk update --no-cache && apk add --no-cache tzdata ca-certificates

# 移动到工作目录
WORKDIR /app

# 优化 Docker 缓存：
# 1. 仅复制 go.mod 和 go.sum
COPY go.mod go.sum ./
# 2. 下载依赖 (需要 ca-certificates)
RUN go mod download

# 3. 复制所有项目源代码
COPY . .

# 4. 编译你指定的 user-api 服务
#    -ldflags="-s -w" 移除调试信息，减小二进制文件大小
#    -o /user-api 指定输出到根目录，方便 Stage 2 复制
RUN go build -ldflags="-s -w" -o /user-api ./app/user/api/

# ----- Stage 2: Final -----
# 使用 scratch 镜像，这是一个完全空白的镜像，体积最小
FROM scratch

# 设置工作目录
WORKDIR /app

# 从 builder 镜像中拷贝必要的系统文件：
# 1. 时区文件 (供 MySQL "loc=" 使用)
# ⭐ 修正：使用 --from=builder
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai

# ⭐ 2. 根证书 (已按要求移除)

# 设置时区环境变量 (供 MySQL "loc=" 使用)
ENV TZ=Asia/Shanghai

# 从 builder 镜像中拷贝编译好的二进制文件
# ⭐ 修正：使用 --from=builder
COPY --from=builder /user-api .

# 从 builder 镜像中拷贝 user-api 所需的 etc 配置目录
# (源路径是 /app (WORKDIR) + /app/user/api/etc (项目路径))
# ⭐ 修正：使用 --from=builder
COPY --from=builder /app/app/user/api/etc/ ./etc/

# 暴露服务端口 (来自 user.yaml)
EXPOSE 8888

# 运行命令, 并显式指定配置文件路径
ENTRYPOINT ["./user-api", "-f", "etc/user.yaml"]