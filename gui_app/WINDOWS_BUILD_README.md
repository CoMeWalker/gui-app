# Windows打包指南

由于Fyne GUI框架的特殊性，在macOS上交叉编译到Windows会有一些挑战。以下是几种解决方案：

## 🎯 推荐方案：在Windows上直接编译

### 步骤：

1. **在Windows系统上安装Go**：
   - 下载并安装Go 1.20+：https://golang.org/dl/

2. **克隆或复制代码到Windows**：
   ```bash
   # 将整个gui_app目录复制到Windows系统
   ```

3. **在Windows上编译**：
   ```cmd
   cd gui_app
   go mod tidy
   go build -o 职业积分管理系统.exe main.go
   ```

4. **运行程序**：
   - 双击 `职业积分管理系统.exe` 即可运行

## 🐳 方案二：使用Docker交叉编译（高级）

### 前提：
- 安装Docker：https://www.docker.com/get-started

### 步骤：

```bash
# 1. 构建Docker镜像
docker build -t fyne-windows-builder .

# 2. 运行容器并复制编译结果
docker run --rm -v $(pwd):/output fyne-windows-builder cp /职业积分管理系统.exe /output/

# 或者直接运行构建
docker run --rm -v $(pwd):/app fyne-windows-builder sh -c "cd /app && go build -o 职业积分管理系统.exe main.go && cp 职业积分管理系统.exe /output/"
```

## 📦 方案三：使用GitHub Actions自动构建

创建一个 `.github/workflows/build.yml` 文件：

```yaml
name: Build Windows Executable

on:
  push:
    branches: [ main ]

jobs:
  build:
    runs-on: windows-latest

    steps:
    - uses: actions/checkout@v2

    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: 1.21

    - name: Build
      run: go build -o 职业积分管理系统.exe main.go

    - name: Upload artifact
      uses: actions/upload-artifact@v2
      with:
        name: windows-executable
        path: 职业积分管理系统.exe
```

## 🔧 故障排除

### 如果编译失败：

1. **检查Go版本**：
   ```bash
   go version
   # 应该显示 go version go1.20.x 或更高
   ```

2. **清理模块缓存**：
   ```bash
   go clean -modcache
   go mod tidy
   ```

3. **检查依赖**：
   ```bash
   go mod verify
   ```

### Windows特定问题：

- **缺少MSYS2或MinGW**：某些情况下可能需要安装MinGW-w64
- **权限问题**：确保有足够的权限读取/写入文件

## 📋 系统要求

### 运行环境：
- **操作系统**：Windows 7 SP1 或更高版本
- **架构**：x64 或 x86
- **内存**：至少256MB可用RAM

### 开发环境：
- **Go版本**：1.20 或更高
- **磁盘空间**：至少100MB可用空间

## 🚀 快速测试

如果只是想测试功能，可以创建一个简单的命令行版本：

```go
// 在main函数中临时替换为：
func main() {
    fmt.Println("青衫似故人 最帅，对吗？")
    fmt.Println("职业积分管理系统已启动！")
}
```

然后编译：
```bash
go build -o test.exe main.go
```

这样可以验证代码逻辑是否正确，然后再处理GUI编译问题。

---

**推荐**：最简单的方法是在Windows系统上直接编译，这样可以避免所有的交叉编译问题。
