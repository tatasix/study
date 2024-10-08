package greedy

import "fmt"

//在柠檬水摊上，每一杯柠檬水的售价为 5 美元。顾客排队购买你的产品，（按账单 bills 支付的顺序）一次购买一杯。
//每位顾客只买一杯柠檬水，然后向你付 5 美元、10 美元或 20 美元。你必须给每个顾客正确找零，也就是说净交易是每位顾客向你支付 5 美元。
//注意，一开始你手头没有任何零钱。
//给你一个整数数组 bills ，其中 bills[i] 是第 i 位顾客付的账。如果你能给每位顾客正确找零，返回 true ，否则返回 false 。
// 示例 1： 输入：bills = [5,5,5,10,20] 输出：true
// 示例 2： 输入：bills = [5,5,10,10,20] 输出：false

func lemonadeChange(bills []int) bool {
	var five int
	var ten int

	for i := 0; i < len(bills); i++ {
		fee := bills[i]
		switch fee {
		case 5:
			five++
		case 10:
			ten++
			five--
		case 20:
			if ten > 0 {
				five--
				ten--
			} else {
				five -= 3
			}

		}
		if five < 0 || ten < 0 {
			return false
		}
	}

	return true
}

func Handle13() {
	fmt.Println(lemonadeChange([]int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}))
}
