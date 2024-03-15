package dynamicplan

import (
	"fmt"
)

// 给定一个整数数组prices，其中第prices[i] 表示第 i 天的股票价格。
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 思路：
// 状态一：持有股票状态（今天买入股票，或者是之前就买入了股票然后没有操作，一直持有）
// 不持有股票状态，这里就有两种卖出股票状态
//		状态二：保持卖出股票的状态（两天前就卖出了股票，度过一天冷冻期。或者是前一天就是卖出股票状态，一直没操作）
//		状态三：今天卖出股票
// 状态四：今天为冷冻期状态，但冷冻期状态不可持续，只有一天！

// 确定递推公式
// 1. 达到买入股票状态（状态一）即：dp[i][0]，有两个具体操作：
// 操作一：前一天就是持有股票状态（状态一），dp[i][0] = dp[i - 1][0]
// 操作二：今天买入了，有两种情况
//
//	前一天是冷冻期（状态四），dp[i - 1][3] - prices[i]
//	前一天是保持卖出股票的状态（状态二），dp[i - 1][1] - prices[i]
//	状态三为什么不行？因为昨天是卖出股票状态，今天就是冷冻期了，就不能再买入了
//
// 那么dp[i][0] = max(dp[i - 1][0], dp[i - 1][3] - prices[i], dp[i - 1][1] - prices[i]);
//
// 2. 达到保持卖出股票状态（状态二）即：dp[i][1]，有两个具体操作：
// 操作一：前一天就是状态二
// 操作二：前一天是冷冻期（状态四）
// dp[i][1] = max(dp[i - 1][1], dp[i - 1][3]);
//
// 3. 达到今天就卖出股票状态（状态三），即：dp[i][2] ，只有一个操作：
// 昨天一定是持有股票状态（状态一），今天卖出
// 即：dp[i][2] = dp[i - 1][0] + prices[i];
//
// 4. 达到冷冻期状态（状态四），即：dp[i][3]，只有一个操作：
// 昨天卖出了股票（状态三）
// dp[i][3] = dp[i - 1][2];
func maxProfit5(prices []int) int {
	l := len(prices)
	dp := make([][]int, l)
	for i, _ := range dp {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0] //持有
	dp[0][1] = 0          //保持之前没有持有的状态
	dp[0][2] = 0          //今天卖出
	dp[0][3] = 0          //冷冻期

	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][1], dp[i-1][3])-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}

	return max(dp[l-1][1], max(dp[l-1][2], dp[l-1][3]))
}

func Handle37() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit5(prices))
}
