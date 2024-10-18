package dynamicplan

// 给你一个整数n，求恰由n个节点组成且节点值从1到n互不相同的二叉搜索树有多少种？
// 返回满足题意的二叉搜索树的种数。
// 输入：n = 3 输出：5
// 输入：n = 1 输出：1
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i < n+1; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
