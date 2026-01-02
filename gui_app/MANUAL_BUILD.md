# 🔧 手动构建Windows程序

## 如果GitHub Actions不工作，使用这个方法

### 📋 方法1：找一台Windows电脑

1. **下载Go**：
   - 访问：https://golang.org/dl/
   - 下载 Windows版本
   - 安装Go 1.21+

2. **下载项目代码**：
   ```cmd
   # 在Windows命令提示符中
   git clone https://github.com/CoMeWalker/gui-app.git
   cd gui-app/gui_app
   ```

3. **安装依赖并构建**：
   ```cmd
   go mod tidy
   go build -o 职业积分管理系统.exe main.go
   ```

4. **运行程序**：
   - 双击 `职业积分管理系统.exe`

### 📋 方法2：使用在线Go Playground

如果没有Windows电脑，可以使用：
- https://play.golang.org/ (但不支持GUI)
- 或者使用Replit、Codesandbox等在线IDE

### 📋 方法3：使用Docker（如果有Docker环境）

```bash
# 构建Docker镜像
docker build -t gui-app-builder .

# 运行构建
docker run --rm -v $(pwd):/output gui-app-builder cp /app/职业积分管理系统.exe /output/
```

### 🎯 验证构建成功

成功的标志：
- 生成 `职业积分管理系统.exe` 文件
- 文件大小约 10-20MB
- 双击可以运行，显示欢迎对话框

### 💡 程序功能

✅ **欢迎界面**：显示"青衫似故人 最帅，对吗？"
✅ **6个职业管理**：完整的GUI界面
✅ **正确的汇总计算**：盈亏 = 本职业积分 × (职业总数-1) - 其他职业积分总和
✅ **复制功能**：一键复制到剪贴板

---

如果GitHub Actions仍然不工作，就用上面的方法手动构建吧！
