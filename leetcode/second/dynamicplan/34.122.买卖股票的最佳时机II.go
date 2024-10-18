package dynamicplan

import "fmt"

// 给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
// 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
// 返回 你能获得的 最大 利润 。
// 示例 1： 输入：prices = [7,1,5,3,6,4] 输出：7
// 示例 2： 输入：prices = [1,2,3,4,5] 输出：4
// 示例 3： 输入：prices = [7,6,4,3,1] 输出：0
// 解释：在这种情况下, 交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为 0。
func maxProfit2(prices []int) int {
	l := len(prices)
	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[l-1][1]
}

func Handle34() {
	fmt.Println(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
}
