package dynamicplan

import "fmt"

// 给定一个字符串 s ，找到其中最长的回文子序列，并返回该序列的长度。可以假设 s 的最大长度为 1000 。
// 示例 1: 输入: "bbbab" 输出: 4 一个可能的最长回文子序列为 "bbbb"。
// 示例 2: 输入:"cbbd" 输出: 2 一个可能的最长回文子序列为 "bb"。
// 回文子串是要连续的，回文子序列可不是连续的！
func longestPalindromeSubseq(s string) int {
	l := len(s)
	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, l)
		for j := range dp[i] {
			if i == j {
				dp[i][i] = 1
			}
		}
	}
	for i := l - 1; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][l-1]
}
func Handle() {
	s := "bbbab"
	fmt.Println(longestPalindromeSubseq(s))
}
