package dynamicplan

import (
	"fmt"
	"math"
)

//给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
//给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
//完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
//示例 1： 输入：n = 12 输出：3 解释：12 = 4 + 4 + 4
//示例 2： 输入：n = 13 输出：2 解释：13 = 4 + 9

func numSquares1(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
		fmt.Println(dp)
	}

	return dp[n]
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		a := i * i
		if a > n {
			break
		}
		for j := a; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-a]+1)
		}
	}
	return dp[n]
}
func Handle24() {
	fmt.Println(numSquares(12))
}
