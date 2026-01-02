# 🔧 GitHub Actions 故障排除指南

## 问题：看不到 "Build Windows Executable" 工作流

### ✅ 检查步骤

#### 1. 确认文件已上传
在你的GitHub仓库中检查以下文件是否存在：

```
.github/
  workflows/
    build-windows.yml
```

#### 2. 检查文件内容
点击 `build-windows.yml` 文件，确保内容完整且正确。

#### 3. 检查Actions标签页
- 确保你在仓库的主页
- 点击 **"Actions"** 标签
- 如果没有看到任何工作流，可能需要等待或重新推送

### 🚀 解决方案

#### **方案1：重新推送代码**
```bash
cd /Users/walker/computer/PersonalWork/code/lsCoder/gui_app

# 强制推送所有文件
git add .
git commit -m "Add GitHub Actions workflow"
git push origin main --force
```

#### **方案2：手动触发工作流**
如果Actions标签页是空的，尝试：
1. 访问你的GitHub仓库
2. 点击 **"Actions"** 标签
3. 如果看到工作流但没有运行，点击 **"Run workflow"**

#### **方案3：检查分支名称**
工作流监听 `main` 和 `master` 分支。确认你的默认分支是这两个之一：
```bash
git branch -a  # 查看所有分支
git branch -m master main  # 如果需要重命名分支
```

#### **方案4：验证工作流语法**
访问：https://github.com/YOUR_USERNAME/YOUR_REPO/actions
如果显示语法错误，检查 `.github/workflows/build-windows.yml` 文件。

### 🔍 常见问题

#### **问题1：Actions标签页空白**
**原因**：工作流文件没有正确上传
**解决**：
```bash
# 确保.github目录被上传
git add .github/
git commit -m "Add workflows"
git push
```

#### **问题2：工作流显示但不运行**
**原因**：触发条件不满足
**解决**：
- 确保推送到了 `main` 或 `master` 分支
- 或者手动触发工作流

#### **问题3：构建失败**
**原因**：Go版本或依赖问题
**解决**：检查Actions日志中的具体错误

### 📞 获取帮助

如果问题仍然存在：

1. **检查GitHub状态**：https://www.githubstatus.com/
2. **查看Actions日志**：找到具体的错误信息
3. **告诉我具体情况**：截图或复制错误信息

### 🎯 验证成功

当你看到：
- ✅ Actions标签页有 "Build Windows Executable"
- ✅ 工作流正在运行（黄色圆点）
- ✅ 构建成功（绿色对勾）
- ✅ 可以下载 "windows-executable"

就说明一切正常了！

---

**💡 提示**：GitHub Actions有时需要5-10分钟来识别新的工作流文件，请耐心等待。
