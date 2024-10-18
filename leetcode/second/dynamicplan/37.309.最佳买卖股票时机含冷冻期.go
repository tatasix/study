package dynamicplan

import "fmt"

// 给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
// 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
// 示例: 输入: [1,2,3,0,2] 输出: 3 解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
func maxProfit5(prices []int) int {
	l := len(prices)
	//dp[i][0] 保持卖出
	//dp[i][1] 买入
	//dp[i][2] 刚卖出
	//dp[i][3] 冷冻期
	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	dp[0][3] = 0
	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][3])
		dp[i][1] = max(dp[i-1][1], max(dp[i-1][0], dp[i-1][3])-prices[i])
		dp[i][2] = dp[i-1][1] + prices[i]
		dp[i][3] = dp[i-1][2]
	}
	return max(dp[l-1][0], max(dp[l-1][2], dp[l-1][3]))
}
func Handle37() {
	fmt.Println(maxProfit5([]int{1, 2, 3, 0, 2}))
}
