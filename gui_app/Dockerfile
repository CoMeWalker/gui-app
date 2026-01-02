# 使用官方Go镜像作为构建环境
FROM golang:1.21-alpine AS builder

# 安装必要的系统依赖
RUN apk add --no-cache gcc musl-dev

# 设置工作目录
WORKDIR /app

# 复制go mod文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 交叉编译为Windows 64位可执行文件
RUN CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -o 职业积分管理系统.exe main.go

# 使用最小镜像作为运行环境
FROM scratch

# 从构建阶段复制可执行文件
COPY --from=builder /app/职业积分管理系统.exe /

# 说明这是什么
LABEL description="职业积分管理系统 - Windows版本"

# 默认命令
CMD ["/职业积分管理系统.exe"]
