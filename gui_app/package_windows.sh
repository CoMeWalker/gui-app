#!/bin/bash

# Windows交叉编译脚本

echo "开始为Windows打包职业积分管理系统..."

# 确保在正确的目录
cd "$(dirname "$0")"

# 64位Windows版本
echo "编译64位Windows版本..."
if GOOS=windows GOARCH=amd64 go build -o "职业积分管理系统_x64.exe" main.go; then
    echo "✅ 64位版本编译成功: 职业积分管理系统_x64.exe"
else
    echo "❌ 64位版本编译失败"
fi

# 32位Windows版本
echo "编译32位Windows版本..."
if GOOS=windows GOARCH=386 go build -o "职业积分管理系统_x86.exe" main.go; then
    echo "✅ 32位版本编译成功: 职业积分管理系统_x86.exe"
else
    echo "❌ 32位版本编译失败"
fi

echo "打包完成！"
