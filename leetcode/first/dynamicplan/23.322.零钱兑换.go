package dynamicplan

import (
	"fmt"
	"math"
)

//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
//你可以认为每种硬币的数量是无限的。
//示例 1： 输入：coins = [1, 2, 5], amount = 11 输出：3
//解释：11 = 5 + 5 + 1
//示例 2： 输入：coins = [2], amount = 3 输出：-1
//示例 3： 输入：coins = [1], amount = 0 输出：0
//示例 4： 输入：coins = [1], amount = 1 输出：1
//示例 5： 输入：coins = [1], amount = 2 输出：2

func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		dp[i] = math.MaxInt32
	}
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			if dp[j-coin] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-coin]+1)
			}

		}
		fmt.Println(dp)
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func Handle23() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i < amount+1; i++ {
		dp[i] = math.MaxInt32
	}
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			if dp[j-coin] == math.MaxInt32 {
				continue
			}
			dp[j] = min(dp[j], dp[j-coin]+1)
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
