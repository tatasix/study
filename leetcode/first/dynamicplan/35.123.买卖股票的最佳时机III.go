package dynamicplan

import (
	"fmt"
)

func maxProfit3(prices []int) int {
	l := len(prices)
	dp := make([][]int, l)
	for i, _ := range dp {
		dp[i] = make([]int, 5)
	}
	dp[0][0], dp[0][1], dp[0][2], dp[0][3], dp[0][4] = 0, -prices[0], 0, -prices[0], 0

	for i := 1; i <= l; i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}

	return dp[l-1][4]
}

func Handle35() {
	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit3(prices))
}
