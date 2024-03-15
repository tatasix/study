package dynamicplan

func minCostClimbingStairs(cost []int) int {
	n := len(cost)

	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 0

	for i := 2; i < n+1; i++ {
		dp[i] = Min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n-1]
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

//给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。
//
//你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
//
//请你计算并返回达到楼梯顶部的最低花费。
