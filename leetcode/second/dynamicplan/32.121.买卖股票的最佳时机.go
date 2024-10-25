package dynamicplan

import (
	"fmt"
	"math"
)

// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
// 示例 1：输入：[7,1,5,3,6,4] 输出：5
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
// 示例 2： 输入：prices = [7,6,4,3,1] 输出：0
func maxProfit11(prices []int) int {
	l := len(prices)
	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, 2)
		dp[i][0] = math.MinInt32
	}
	dp[0][0] = 0 - prices[0]
	dp[0][1] = 0

	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], 0-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[l-1][1]
}
func maxProfit12(prices []int) int {
	minV := math.MaxInt32
	var result int
	for i := 0; i < len(prices); i++ {
		if prices[i] < minV {
			minV = prices[i]
		}
		if result < prices[i]-minV {
			result = prices[i] - minV
		}
	}
	return result
}
func maxProfit1(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}
	dp := make([][]int, l)
	for i := range dp {
		// 0 持有，1 不持有
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], 0-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[l-1][1]
}
func Handle32() {
	fmt.Println(maxProfit1([]int{7, 6, 4, 3, 1}))
}
