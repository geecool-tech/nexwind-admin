# 基础镜像
FROM registry.cn-chengdu.aliyuncs.com/nexwind_backend/golang_alpine_1_21:latest AS builder

# 设置工作目录
WORKDIR /app
ENV GOPROXY=https://goproxy.cn,direct

# 复制go.mod和go.sum以下载依赖
COPY go.mod .
COPY go.sum .

# 下载依赖
RUN go mod download

# 将源代码复制到工作目录
COPY . .

# 编译应用程序
RUN go build -o ./cmd/v1/main ./cmd/v1

# 新的镜像阶段
FROM registry.cn-chengdu.aliyuncs.com/nexwind_backend/alpine:latest

# 安装运行所需的依赖
# RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制编译后的可执行文件到新的镜像中
COPY --from=builder /app/cmd/v1/main ./cmd/v1/main
# 复制配置文件到镜像中
COPY --from=builder /app/config/v1/app.yaml .
#COPY --from=builder /frontend/logs .

# 设置环境变量
ENV ENVIRONMENT=production

# 运行应用程序
CMD ["./cmd/v1/main"]
