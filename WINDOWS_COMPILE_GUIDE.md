# Windows编译指南

## 📋 在Windows上编译职业积分管理系统

### 前提条件

1. **安装Go语言**
   - 下载地址：https://golang.org/dl/
   - 选择 `go1.21.x.windows-amd64.msi`
   - 安装时选择默认路径

2. **安装Git** (可选，用于克隆代码)
   - 下载地址：https://git-scm.com/downloads
   - 安装时选择默认选项

3. **验证安装**
   ```cmd
   go version
   # 应该显示: go version go1.21.x windows/amd64

   git version
   # 应该显示: git version 2.x.x.windows.x
   ```

### 📥 获取源代码

#### 方法1：从GitHub克隆
```cmd
git clone https://github.com/CoMeWalker/gui-app.git
cd gui-app
```

#### 方法2：下载ZIP包
1. 访问：https://github.com/CoMeWalker/gui-app
2. 点击 **"Code"** → **"Download ZIP"**
3. 解压到本地文件夹

### 🔧 编译步骤

#### 1. 进入项目目录
```cmd
cd gui-app  # 如果是从ZIP下载的
# 或者已经在克隆的目录中
```

#### 2. 下载依赖
```cmd
go mod tidy
```

#### 3. 编译程序
```cmd
go build -o 职业积分管理系统.exe main.go
```

#### 4. 验证编译结果
```cmd
dir 职业积分管理系统.exe
# 应该看到生成的可执行文件
```

### 🎯 完整命令序列

```cmd
# 1. 克隆项目
git clone https://github.com/CoMeWalker/gui-app.git
cd gui-app

# 2. 下载依赖
go mod tidy

# 3. 编译程序
go build -o 职业积分管理系统.exe main.go

# 4. 运行程序（可选）
职业积分管理系统.exe
```

### 🛠️ 故障排除

#### 问题1：找不到go命令
```
'go' is not recognized as an internal or external command
```
**解决方法**：
1. 检查Go是否正确安装
2. 重启命令提示符
3. 或者使用完整路径：`"C:\Program Files\Go\bin\go.exe" build main.go`

#### 问题2：编译失败
```
package fyne.io/fyne/v2/app: cannot find package
```
**解决方法**：
```cmd
# 清理模块缓存
go clean -modcache

# 重新下载依赖
go mod download

# 再次尝试
go build -o 职业积分管理系统.exe main.go
```

#### 问题3：权限问题
```
Access is denied
```
**解决方法**：
- 以管理员身份运行命令提示符
- 或者保存到其他目录：`go build -o C:\Users\YourName\Desktop\职业积分管理系统.exe main.go`

### 🎨 程序功能说明

编译成功后，你将获得一个完整的职业积分管理系统：

✅ **欢迎界面**：显示个性化提示
✅ **6个职业管理**：青城QC、仙禽XQ、百花HH、和尚HS、天净TJ、峨眉EM
✅ **玩家姓名输入**：文本框输入玩家姓名
✅ **积分计数器**：+/-按钮调整积分数值
✅ **智能汇总计算**：正确的盈亏计算公式
✅ **一键复制**：复制到剪贴板
✅ **临时提示**：操作反馈

### 📂 文件结构

编译成功后，你的项目目录应该包含：
```
gui-app/
├── main.go                 # 主程序
├── go.mod                  # Go模块文件
├── go.sum                  # 依赖校验
├── 职业积分管理系统.exe     # 生成的可执行文件 ⭐
├── README.md              # 说明文档
└── .github/               # GitHub Actions配置
```

### 🚀 运行程序

双击 `职业积分管理系统.exe` 文件即可运行！

---

## 🎊 成功标志

如果看到：
- ✅ 生成了 `职业积分管理系统.exe` 文件
- ✅ 双击运行后显示欢迎对话框
- ✅ 可以正常使用所有功能

就说明编译完全成功了！

有任何编译问题随时告诉我！🎉
