// 简化的命令行版本 - 可以在任何系统上编译
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Profession struct {
	name       string
	playerName string
	score      float64
}

func main() {
	fmt.Println("=== 职业积分管理系统 ===")
	fmt.Println("青衫似故人 最帅，对吗？(y/n)")

	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))

	if response != "y" && response != "yes" {
		fmt.Println("再见！")
		return
	}

	fmt.Println("欢迎使用职业积分管理系统！")
	fmt.Println()

	professions := []Profession{
		{name: "青城QC"},
		{name: "仙禽XQ"},
		{name: "百花HH"},
		{name: "和尚HS"},
		{name: "天净TJ"},
		{name: "峨眉EM"},
	}

	// 输入玩家姓名和积分
	for i := range professions {
		for {
			fmt.Printf("请输入%s的玩家姓名: ", professions[i].name)
			playerName, _ := reader.ReadString('\n')
			playerName = strings.TrimSpace(playerName)

			if playerName == "" {
				fmt.Printf("%s: 职业与玩家未绑定\n", professions[i].name)
				break
			}

			professions[i].playerName = playerName

			fmt.Printf("请输入%s的积分数值 (默认0): ", professions[i].name)
			scoreInput, _ := reader.ReadString('\n')
			scoreInput = strings.TrimSpace(scoreInput)

			if scoreInput != "" {
				if score, err := strconv.ParseFloat(scoreInput, 64); err == nil {
					professions[i].score = score
				}
			}
			break
		}
	}

	// 输入兑换比例
	var rate float64 = 1.0
	for {
		fmt.Print("请输入积分兑换金币的比例 (默认1.0): ")
		rateInput, _ := reader.ReadString('\n')
		rateInput = strings.TrimSpace(rateInput)

		if rateInput == "" {
			break
		}

		if r, err := strconv.ParseFloat(rateInput, 64); err == nil {
			rate = r
			break
		}
		fmt.Println("请输入有效的数字！")
	}

	fmt.Println("\n=== 汇总结果 ===")

	// 计算所有职业积分的总和
	totalAllScores := 0.0
	validProfessionCount := 0
	for _, prof := range professions {
		if prof.score > 0 {
			totalAllScores += prof.score
			validProfessionCount++
		}
	}

	// 显示结果
	for _, prof := range professions {
		if prof.playerName == "" {
			fmt.Printf("%s: 职业与玩家未绑定\n", prof.name)
		} else {
			// 计算：本职业积分 × (职业总数-1) - 其他职业积分总和
			otherScoresTotal := totalAllScores - prof.score
			profit := prof.score*float64(validProfessionCount-1) - otherScoresTotal
			goldProfit := profit * rate

			profitStr := fmt.Sprintf("%.2f", profit)
			if profit > 0 {
				profitStr = "+" + profitStr
			}

			goldProfitStr := fmt.Sprintf("%.2f", goldProfit)
			if goldProfit > 0 {
				goldProfitStr = "+" + goldProfitStr
			}

			fmt.Printf("%s (%s): 积分 %.2f - 其他职业总和 %.2f = 积分盈亏 %s, 金币盈亏 %s\n",
				prof.name, prof.playerName, prof.score, otherScoresTotal, profitStr, goldProfitStr)
		}
	}

	// 复制到剪贴板（模拟）
	var results []string
	for _, prof := range professions {
		if prof.playerName != "" {
			results = append(results, fmt.Sprintf("%s[%s] %.2f", prof.playerName, prof.name, prof.score))
		}
	}

	if len(results) > 0 {
		copyText := strings.Join(results, " | ")
		fmt.Println("\n复制到剪贴板的格式:")
		fmt.Println(copyText)
		fmt.Println("\n✅ 数据已准备好复制！")
	}

	fmt.Println("\n按Enter键退出...")
	reader.ReadString('\n')
}
