package dynamicplan

import (
	"fmt"
)

// 给你一个整数数组 prices 和一个整数 k ，其中 prices[i] 是某支给定的股票在第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成k笔交易。也就是说,你最多可以买 k 次,卖 k 次。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
func maxProfit4(k int, prices []int) int {
	l := len(prices)
	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, 2*k+1)
	}
	for j := 1; j <= k; j++ {
		n := j*2 - 1
		dp[0][n] = 0 - prices[0]
	}
	for i := 1; i < l; i++ {
		dp[i][0] = dp[i-1][0]
		for j := 1; j <= 2*k; j = j + 2 {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]+prices[i])
		}
	}
	return dp[l-1][2*k]
}

func Handle36() {
	k := 2
	prices := []int{3, 2, 6, 5, 0, 3}
	fmt.Println(maxProfit4(k, prices))
}
