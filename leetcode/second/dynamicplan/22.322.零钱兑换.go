package dynamicplan

import (
	"fmt"
	"math"
)

//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的
//硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。 你可以认为每种硬币的数量是无限的。
//示例 1: 输入:coins = [1, 2, 5], amount = 11 输出:3 解释:11 = 5 + 5 + 1
//示例 2: 输入:coins = [2], amount = 3 输出:-1
//示例 3: 输入:coins = [1], amount = 0 输出:0
//示例 4: 输入:coins = [1], amount = 1 输出:1
//示例 5: 输入:coins = [1], amount = 2 输出:2

func coinChange1(coins []int, amount int) int {
	l := len(coins)
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i < l; i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
		fmt.Println(dp)
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
func coinChange(coins []int, amount int) int {
	l := len(coins)
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i < l; i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	return dp[amount]
}
func Handle22() {
	//fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
func productExceptSelf(nums []int) []int {
	l := len(nums)
	left := make([]int, l)
	right := make([]int, l)
	left[0] = 1
	right[l-1] = 1
	for i := 1; i < l; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	for i := l - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	for i := 0; i < l; i++ {
		left[i] *= right[i]
	}
	return left
}
