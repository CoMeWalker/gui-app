// Fyne布局示例 - 展示不同的位置控制方式
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showLayoutExamples() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Fyne布局示例")

	// 1. HBox - 水平排列（常用）
	hbox := container.NewHBox(
		widget.NewButton("左边", nil),
		widget.NewButton("中间", nil),
		widget.NewButton("右边", nil),
	)

	// 2. VBox - 垂直排列
	vbox := container.NewVBox(
		widget.NewButton("顶部", nil),
		widget.NewButton("中间", nil),
		widget.NewButton("底部", nil),
	)

	// 3. Center - 居中对齐
	center := container.NewCenter(
		widget.NewButton("居中按钮", nil),
	)

	// 4. Border - 边框布局（精确位置控制）
	border := container.NewBorder(
		widget.NewButton("顶部", nil), // 上
		widget.NewButton("底部", nil), // 下
		widget.NewButton("左侧", nil), // 左
		widget.NewButton("右侧", nil), // 右
		widget.NewLabel("中间区域"),     // 中间
	)

	// 5. Grid - 网格布局
	grid := container.NewGridWithColumns(3,
		widget.NewButton("1", nil),
		widget.NewButton("2", nil),
		widget.NewButton("3", nil),
		widget.NewButton("4", nil),
		widget.NewButton("5", nil),
		widget.NewButton("6", nil),
	)

	// 6. Padded - 添加内边距
	padded := container.NewPadded(
		widget.NewButton("带内边距的按钮", nil),
	)

	// 组合使用示例
	complex := container.NewVBox(
		widget.NewLabel("复杂布局示例:"),
		container.NewBorder(
			nil,
			container.NewHBox(widget.NewButton("确定", nil), widget.NewButton("取消", nil)),
			nil,
			nil,
			container.NewCenter(widget.NewLabel("主要内容区域")),
		),
	)

	mainContent := container.NewVBox(
		widget.NewLabel("Fyne布局方式演示:"),
		widget.NewSeparator(),
		container.NewVBox(
			widget.NewLabel("1. 水平排列 (HBox):"), hbox,
			widget.NewLabel("2. 垂直排列 (VBox):"), vbox,
			widget.NewLabel("3. 居中对齐 (Center):"), center,
			widget.NewLabel("4. 边框布局 (Border):"), border,
			widget.NewLabel("5. 网格布局 (Grid):"), grid,
			widget.NewLabel("6. 内边距 (Padded):"), padded,
			widget.NewLabel("7. 复杂组合:"), complex,
		),
	)

	myWindow.SetContent(container.NewScroll(mainContent))
	myWindow.Resize(fyne.NewSize(600, 800))
	myWindow.ShowAndRun()
}
