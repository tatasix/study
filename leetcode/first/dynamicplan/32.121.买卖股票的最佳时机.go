package dynamicplan

import (
	"fmt"
	"math"
)

//给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
//你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

func maxProfit1(prices []int) int {
	var result int
	minPrice := math.MaxInt32
	for i := 0; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		result = max(result, prices[i]-minPrice)
	}
	return result
}

func maxProfit11(prices []int) int {
	l := len(prices)
	dp := make([][]int, l)
	for i, _ := range prices {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0 - prices[0]
	dp[0][1] = 0
	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], 0-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return max(dp[l-1][0], dp[l-1][1])
}

func Handle32() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit11(prices))
}
