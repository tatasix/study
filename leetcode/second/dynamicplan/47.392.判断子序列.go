package dynamicplan

import "fmt"

// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
// 示例 1： 输入：s = "abc", t = "ahbgdc" 输出：true
// 示例 2： 输入：s = "axc", t = "ahbgdc" 输出：false
func isSubsequence1(s, t string) bool {
	ls, lt := len(s), len(t)
	dp := make([][]int, ls+1)
	var maxLength int
	for i := range dp {
		dp[i] = make([]int, lt+1)
	}
	for i := 1; i <= ls; i++ {
		for j := 1; j <= lt; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
			if maxLength < dp[i][j] {
				maxLength = dp[i][j]
			}
		}
	}
	return maxLength == ls
}

func isSubsequence(s, t string) bool {
	l1, l2 := len(s), len(t)
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}
	result := 0
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
			if dp[i][j] > result {
				result = dp[i][j]
			}
		}
	}

	return result == l1
}

func Handle47() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
