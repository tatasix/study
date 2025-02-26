package dynamicplan

import "fmt"

//数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i](下标从 0 开始)。
//每当你爬上一个阶梯你都要花费对应的体力值，一旦支付了相应的体力值，你就可以选择向上爬一个阶梯或者爬两个阶梯。
//请你找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。
//示例 1: 输入:cost = [10, 15, 20] 输出:15
//示例 2: 输入:cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1] 输出:6

func minCostClimbingStairs(cost []int) int {
	l := len(cost)
	dp := make([]int, l+1)
	dp[0] = 0
	dp[1] = 0
	if l < 2 {
		return 0
	}
	for i := 2; i < l+1; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[l]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Handle4() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}
