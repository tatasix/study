package dynamicplan

import "fmt"

// 给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
// 字符串的一个 子序列 是指，通过删除一些(也可以不删除)字符且不干扰剩余字符相对位置所组成的新 字符串。(例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是)
// 题目数据保证答案符合 32 位带符号整数范围。
// 示例 1： 输入：s = "rabbbit", t = "rabbit" 输出：3
// 如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
// 示例 2： 输入：s = "babgbag", t = "bag" 输出：5
// 如下所示, 有 5 种可以从 s 中得到 "bag" 的方案。
func numDistinct1(s string, t string) int {
	ls, lt := len(s), len(t)
	dp := make([][]int, ls+1)
	for i := range dp {
		dp[i] = make([]int, lt+1)
		dp[i][0] = 1
	}
	for i := 1; i <= ls; i++ {
		for j := 1; j <= lt; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[ls][lt]
}

func numDistinct(s string, t string) int {
	l1, l2 := len(s), len(t)
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
		dp[i][0] = 1
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[l1][l2]
}

func Handle48() {
	s := "rabbbit"
	t := "rabbit"
	fmt.Println(numDistinct(s, t))
}
