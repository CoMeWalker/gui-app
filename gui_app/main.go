package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// 职业数据结构
type Profession struct {
	name       string
	entry      *widget.Entry
	countLabel *widget.Label
}

// 显示临时消息的函数
func showTemporaryMessage(label *widget.Label, message string, duration time.Duration) {
	label.SetText(message)
	label.Show()

	// 创建定时器，在指定时间后隐藏消息
	time.AfterFunc(duration, func() {
		label.Hide()
	})
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("职业积分管理系统")
	myWindow.Resize(fyne.NewSize(800, 600))

	// 先设置窗口内容，然后显示窗口
	setupMainWindow(myWindow)
	myWindow.Show()

	// 创建欢迎对话框
	welcomeDialog := dialog.NewConfirm("欢迎", "青衫似故人 最帅，对吗？", func(confirmed bool) {
		if !confirmed {
			// 用户点击"否"，退出程序
			myApp.Quit()
		}
		// 用户点击"是"，继续使用程序（窗口已经显示）
	}, myWindow)

	// 显示欢迎对话框
	welcomeDialog.Show()

	// 启动应用
	myApp.Run()
}

// 设置主窗口内容的函数
func setupMainWindow(myWindow fyne.Window) {

	// 创建状态标签用于显示临时消息
	statusLabel := widget.NewLabel("")
	statusLabel.Hide()

	// 创建6个职业，每个都有输入框和计数器
	professions := []*Profession{
		{name: "青城QC", entry: widget.NewEntry(), countLabel: widget.NewLabel("0")},
		{name: "仙禽XQ", entry: widget.NewEntry(), countLabel: widget.NewLabel("0")},
		{name: "百花HH", entry: widget.NewEntry(), countLabel: widget.NewLabel("0")},
		{name: "和尚HS", entry: widget.NewEntry(), countLabel: widget.NewLabel("0")},
		{name: "天净TJ", entry: widget.NewEntry(), countLabel: widget.NewLabel("0")},
		{name: "峨眉EM", entry: widget.NewEntry(), countLabel: widget.NewLabel("0")},
	}

	// 设置输入框最小宽度（能显示6个汉字），增加宽度
	for _, prof := range professions {
		prof.entry.Resize(fyne.NewSize(150, 36)) // 从120增加到150
	}

	rateEntry := widget.NewEntry()
	rateEntry.SetText("1.0")
	rateEntry.Resize(fyne.NewSize(200, 36)) // 固定宽度为200像素

	// 结果显示
	resultEntry := widget.NewMultiLineEntry()
	resultEntry.Disable()
	resultEntry.Resize(fyne.NewSize(800, 600)) // 高度增加300像素

	// 创建可滚动的容器
	scrollContainer := container.NewScroll(resultEntry)
	scrollContainer.SetMinSize(fyne.NewSize(800, 300))

	// 创建职业区块的函数
	createProfessionBlock := func(prof *Profession) *fyne.Container {
		increaseBtn := widget.NewButton("  +  ", func() {
			currentStr := prof.countLabel.Text
			current, _ := strconv.Atoi(currentStr)
			current++
			prof.countLabel.SetText(strconv.Itoa(current))
		})
		increaseBtn.Resize(fyne.NewSize(80, 36)) // 宽度扩大30像素

		decreaseBtn := widget.NewButton("  -  ", func() {
			currentStr := prof.countLabel.Text
			current, _ := strconv.Atoi(currentStr)
			if current > 0 {
				current--
				prof.countLabel.SetText(strconv.Itoa(current))
			}
		})
		decreaseBtn.Resize(fyne.NewSize(80, 36)) // 宽度扩大30像素

		return container.NewVBox(
			widget.NewLabelWithStyle(prof.name, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			container.NewPadded(prof.entry),
			container.NewHBox(
				widget.NewLabel("数量:"),
				prof.countLabel,
				increaseBtn,
				decreaseBtn,
			),
			widget.NewSeparator(), // 每个职业区块底部添加分割线
		)
	}

	// 复制积分按钮
	copyBtn := widget.NewButton("复制积分", func() {
		var parts []string

		for _, prof := range professions {
			text := prof.entry.Text
			count := prof.countLabel.Text
			if text != "" {
				parts = append(parts, fmt.Sprintf("%s[%s] %s", text, prof.name, count))
			}
		}

		result := strings.Join(parts, " | ")
		clipboard := myWindow.Clipboard()
		clipboard.SetContent(result)

		// 显示临时成功消息
		showTemporaryMessage(statusLabel, "✓ 积分已复制到剪贴板", 3*time.Second)
	})

	// 汇总按钮
	summaryBtn := widget.NewButton("汇总计算", func() {
		rateStr := rateEntry.Text
		rate, err := strconv.ParseFloat(rateStr, 64)
		if err != nil {
			showTemporaryMessage(statusLabel, "✗ 请输入有效的兑换比例", 3*time.Second)
			return
		}

		// 首先计算所有职业积分的总和（从计数器获取积分）
		totalAllScores := 0.0
		validProfessionCount := 0

		for _, prof := range professions {
			scoreStr := prof.countLabel.Text // 从计数器获取积分
			if scoreStr != "0" {             // 如果计数器不是0，说明有积分
				score, err := strconv.ParseFloat(scoreStr, 64)
				if err == nil {
					totalAllScores += score
					validProfessionCount++
				}
			}
		}

		var results []string

		for _, prof := range professions {
			playerName := prof.entry.Text    // 玩家姓名
			scoreStr := prof.countLabel.Text // 积分数值

			if playerName == "" {
				results = append(results, fmt.Sprintf("%s: 职业与玩家未绑定", prof.name))
				continue
			}

			score, err := strconv.ParseFloat(scoreStr, 64)
			if err != nil {
				results = append(results, fmt.Sprintf("%s: 积分数值无效", prof.name))
				continue
			}

			// 计算：本职业积分 × (职业总数-1) - 其他职业积分总和
			otherScoresTotal := totalAllScores - score
			profit := score*float64(validProfessionCount-1) - otherScoresTotal
			goldProfit := profit * rate

			profitStr := fmt.Sprintf("%.2f", profit)
			if profit > 0 {
				profitStr = "+" + profitStr
			}

			goldProfitStr := fmt.Sprintf("%.2f", goldProfit)
			if goldProfit > 0 {
				goldProfitStr = "+" + goldProfitStr
			}

			results = append(results, fmt.Sprintf("%s (%s): 积分 %.2f - 其他职业总和 %.2f = 积分盈亏 %s, 金币盈亏 %s",
				prof.name, playerName, score, otherScoresTotal, profitStr, goldProfitStr))
		}

		resultEntry.SetText(strings.Join(results, "\n"))

		// 复制到剪贴板
		copyText := strings.Join(results, " | ")
		clipboard := myWindow.Clipboard()
		clipboard.SetContent(copyText)

		// 显示临时成功消息
		showTemporaryMessage(statusLabel, "✓ 汇总结果已复制到剪贴板", 3*time.Second)
	})

	// 设置按钮大小
	summaryBtn.Resize(fyne.NewSize(100, 40))

	// 界面布局
	// 创建职业区块，增加间距
	var professionBlocks []fyne.CanvasObject
	for _, prof := range professions {
		professionBlocks = append(professionBlocks,
			container.NewPadded(createProfessionBlock(prof))) // 每个职业区块增加内边距
	}

	// 职业区块横向排列
	professionsRow := container.NewHBox(professionBlocks...)

	// 积分兑换区域 - 标签换行显示
	rateControls := container.NewVBox(
		widget.NewLabel("积分兑换金币比例:"),
		container.NewHBox(
			container.NewPadded(rateEntry),
			container.NewPadded(summaryBtn),
		),
	)

	// 功能按钮区域 - 使用Border布局保持控件原始大小并居中
	buttonsRow := container.NewBorder(
		nil,                          // 顶部
		nil,                          // 底部
		container.NewPadded(copyBtn), // 左侧：复制积分按钮
		nil,                          // 右侧
		container.NewCenter(container.NewPadded(rateControls)), // 中间居中：积分兑换区域
	)

	// 状态消息区域
	statusRow := container.NewHBox(statusLabel)

	// 主布局 - 增加行间距
	content := container.NewVBox(
		container.NewPadded(professionsRow),        // 职业行增加外边距
		container.NewPadded(widget.NewSeparator()), // 分割线增加边距
		container.NewPadded(buttonsRow),            // 按钮行增加较大外边距
		container.NewPadded(scrollContainer),       // 可滚动的结果显示区域
		container.NewPadded(statusRow),             // 状态消息行 - 移到最后
	)

	myWindow.SetContent(content)
}
