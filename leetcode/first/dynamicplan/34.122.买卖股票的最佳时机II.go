package dynamicplan

import (
	"fmt"
)

func maxProfit21(prices []int) int {
	var result int
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			result += prices[i] - prices[i-1]
		}
	}
	return result
}

func maxProfit22(prices []int) int {
	l := len(prices)
	dp := make([][]int, l)
	for i, _ := range prices {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0 - prices[0]
	dp[0][1] = 0
	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return max(dp[l-1][0], dp[l-1][1])
}

func Handle34() {
	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit22(prices))
}
