package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MainWindow struct {
	*walk.MainWindow
	// 输入框
	qcEdit *walk.LineEdit
	xqEdit *walk.LineEdit
	hhEdit *walk.LineEdit
	hsEdit *walk.LineEdit
	tjEdit *walk.LineEdit
	emEdit *walk.LineEdit

	// 数字显示和控制
	numberLabel *walk.Label
	rateEdit    *walk.LineEdit

	// 结果显示
	resultText *walk.TextEdit

	// 状态显示
	statusLabel *walk.Label
}

func (mw *MainWindow) copyScores() {
	var parts []string

	professions := []struct {
		name  string
		edit  *walk.LineEdit
		label *walk.Label
	}{
		{"青城QC", mw.qcEdit, mw.numberLabel},
		{"仙禽XQ", mw.xqEdit, mw.numberLabel},
		{"百花HH", mw.hhEdit, mw.numberLabel},
		{"和尚HS", mw.hsEdit, mw.numberLabel},
		{"天净TJ", mw.tjEdit, mw.numberLabel},
		{"峨眉EM", mw.emEdit, mw.numberLabel},
	}

	for _, prof := range professions {
		playerName, _ := prof.edit.Text()
		score, _ := prof.label.Text()
		if playerName != "" {
			parts = append(parts, fmt.Sprintf("%s[%s] %s", playerName, prof.name, score))
		}
	}

	result := strings.Join(parts, " | ")
	walk.Clipboard().SetText(result)
	mw.showTemporaryMessage("✓ 积分已复制到剪贴板")
}

func (mw *MainWindow) calculateSummary() {
	rateStr, _ := mw.rateEdit.Text()
	rate, err := strconv.ParseFloat(rateStr, 64)
	if err != nil {
		mw.showTemporaryMessage("✗ 请输入有效的兑换比例")
		return
	}

	professions := []struct {
		name  string
		edit  *walk.LineEdit
		label *walk.Label
	}{
		{"青城QC", mw.qcEdit, mw.numberLabel},
		{"仙禽XQ", mw.xqEdit, mw.numberLabel},
		{"百花HH", mw.hhEdit, mw.numberLabel},
		{"和尚HS", mw.hsEdit, mw.numberLabel},
		{"天净TJ", mw.tjEdit, mw.numberLabel},
		{"峨眉EM", mw.emEdit, mw.numberLabel},
	}

	// 首先计算所有职业积分的总和（从计数器获取积分）
	totalAllScores := 0.0
	validProfessionCount := 0

	for _, prof := range professions {
		scoreStr, _ := prof.label.Text() // 从计数器获取积分
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
		playerName, _ := prof.edit.Text() // 玩家姓名
		scoreStr, _ := prof.label.Text()  // 积分数值

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

	mw.resultText.SetText(strings.Join(results, "\n"))

	copyText := strings.Join(results, " | ")
	walk.Clipboard().SetText(copyText)
	mw.showTemporaryMessage("✓ 汇总结果已复制到剪贴板")
}

func (mw *MainWindow) increaseNumber() {
	currentStr := mw.numberLabel.Text()
	current, _ := strconv.Atoi(currentStr)
	current++
	mw.numberLabel.SetText(strconv.Itoa(current))
}

func (mw *MainWindow) decreaseNumber() {
	currentStr := mw.numberLabel.Text()
	current, _ := strconv.Atoi(currentStr)
	if current > 0 {
		current--
		mw.numberLabel.SetText(strconv.Itoa(current))
	}
}

func (mw *MainWindow) showTemporaryMessage(message string) {
	mw.statusLabel.SetText(message)
	mw.statusLabel.SetVisible(true)

	// 3秒后隐藏消息
	go func() {
		time.Sleep(3 * time.Second)
		mw.statusLabel.SetVisible(false)
	}()
}

func main() {
	var mw MainWindow

	// 创建欢迎对话框
	walk.MsgBox(nil, "欢迎", "青衫似故人 最帅，对吗？", walk.MsgBoxIconQuestion|walk.MsgBoxOKCancel)

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "职业积分管理系统",
		MinSize:  Size{900, 700},
		Layout:   VBox{},
		Children: []Widget{
			// 职业输入区域
			Composite{
				Layout: HBox{},
				Children: []Widget{
					// 青城QC
					Composite{
						Layout: VBox{},
						Children: []Widget{
							Label{Text: "青城QC"},
							LineEdit{AssignTo: &mw.qcEdit},
							PushButton{Text: "+", OnClicked: mw.increaseNumber},
							PushButton{Text: "-", OnClicked: mw.decreaseNumber},
						},
					},
					// 仙禽XQ
					Composite{
						Layout: VBox{},
						Children: []Widget{
							Label{Text: "仙禽XQ"},
							LineEdit{AssignTo: &mw.xqEdit},
							PushButton{Text: "+", OnClicked: mw.increaseNumber},
							PushButton{Text: "-", OnClicked: mw.decreaseNumber},
						},
					},
					// 百花HH
					Composite{
						Layout: VBox{},
						Children: []Widget{
							Label{Text: "百花HH"},
							LineEdit{AssignTo: &mw.hhEdit},
							PushButton{Text: "+", OnClicked: mw.increaseNumber},
							PushButton{Text: "-", OnClicked: mw.decreaseNumber},
						},
					},
					// 和尚HS
					Composite{
						Layout: VBox{},
						Children: []Widget{
							Label{Text: "和尚HS"},
							LineEdit{AssignTo: &mw.hsEdit},
							PushButton{Text: "+", OnClicked: mw.increaseNumber},
							PushButton{Text: "-", OnClicked: mw.decreaseNumber},
						},
					},
					// 天净TJ
					Composite{
						Layout: VBox{},
						Children: []Widget{
							Label{Text: "天净TJ"},
							LineEdit{AssignTo: &mw.tjEdit},
							PushButton{Text: "+", OnClicked: mw.increaseNumber},
							PushButton{Text: "-", OnClicked: mw.decreaseNumber},
						},
					},
					// 峨眉EM
					Composite{
						Layout: VBox{},
						Children: []Widget{
							Label{Text: "峨眉EM"},
							LineEdit{AssignTo: &mw.emEdit},
							PushButton{Text: "+", OnClicked: mw.increaseNumber},
							PushButton{Text: "-", OnClicked: mw.decreaseNumber},
						},
					},
				},
			},

			// 数量显示
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Label{Text: "数量:"},
					Label{AssignTo: &mw.numberLabel, Text: "0"},
				},
			},

			// 分割线
			HSeparator{},

			// 功能按钮区域
			Composite{
				Layout: VBox{},
				Children: []Widget{
					// 第一行：复制积分按钮
					Composite{
						Layout: HBox{},
						Children: []Widget{
							PushButton{
								Text:      "复制积分",
								OnClicked: mw.copyScores,
							},
						},
					},

					// 第二行：兑换比例和汇总按钮
					Composite{
						Layout: HBox{},
						Children: []Widget{
							Label{Text: "积分兑换金币比例:"},
							LineEdit{AssignTo: &mw.rateEdit, Text: "1.0"},
							PushButton{
								Text:      "汇总计算",
								OnClicked: mw.calculateSummary,
							},
						},
					},
				},
			},

			// 状态消息
			Label{AssignTo: &mw.statusLabel, Visible: false},

			// 结果显示区域
			TextEdit{
				AssignTo: &mw.resultText,
				ReadOnly: true,
				MinSize:  Size{800, 300},
			},
		},
	}.Create()); err != nil {
		panic(err)
	}

	mw.Run()
}
