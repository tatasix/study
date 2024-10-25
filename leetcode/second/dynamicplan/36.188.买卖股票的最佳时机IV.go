package dynamicplan

import "fmt"

// 给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 示例 1： 输入：k = 2, prices = [2,4,1] 输出：2
// 示例 2：输入：k = 2, prices = [3,2,6,5,0,3] 输出：7
func maxProfit41(k int, prices []int) int {
	l := len(prices)
	dp := make([][]int, l)
	l1 := 2*k + 1
	for i := range dp {
		dp[i] = make([]int, l1)
	}
	for i := 0; i < l1-1; i += 2 {
		dp[0][i+1] = -prices[0]
	}

	for i := 1; i < l; i++ {
		dp[i][0] = dp[i-1][0]
		for j := 1; j < l1-1; j += 2 {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]+prices[i])
		}
	}
	var res int
	for i := 0; i < l1; i += 2 {
		if res < dp[l-1][i] {
			res = dp[l-1][i]
		}
	}
	return res
}
func maxProfit4(k int, prices []int) int {
	l := len(prices)
	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, 2*k+1)
	}
	for i := 1; i < 2*k; i += 2 {
		dp[0][i] = -prices[0]
	}
	var result int
	for i := 1; i < l; i++ {
		dp[i][0] = dp[i-1][0]
		for j := 1; j < 2*k; j += 2 {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]+prices[i])
			if dp[i][j+1] > result {
				result = dp[i][j+1]
			}
		}
	}
	return result
}
func Handle36() {
	fmt.Println(maxProfit4(2, []int{3, 2, 6, 5, 0, 3}))
}
