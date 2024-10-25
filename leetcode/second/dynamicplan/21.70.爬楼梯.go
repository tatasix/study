package dynamicplan

// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢? 注意:给定 n 是一个正整数。
// 示例 1: 输入: 2 输出: 2 解释: 有两种方法可以爬到楼顶。 1. 1阶+1阶 2. 2阶
// 示例 2: 输入: 3 输出: 3 解释: 有三种方法可以爬到楼顶。 1. 1阶+1阶+1阶 2. 1阶+2阶 3. 2阶+1阶
func climbStairs2(n, m int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := n; j <= m; j++ {
			dp[j] += dp[i-j]
		}
	}

	return 0
}

func climbStairs3(n, m int) int {

	return 0
}
