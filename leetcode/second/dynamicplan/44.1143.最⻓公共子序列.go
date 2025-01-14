package dynamicplan

import "fmt"

//给定两个字符串 text1 和 text2，返回这两个字符串的最⻓公共子序列的⻓度。
//一个字符串的 子序列 是指这样一个新的字符串:它是由原字符串在不改变字符的相对顺序的情况下删除 某些字符(也可以不删除任何字符)后组成的新字符串。
//例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是 这两个字符串所共同拥有的子序列。
//若这两个字符串没有公共子序列，则返回 0。 示例 1:
//输入:text1 = "abcde", text2 = "ace" 输出:3
//解释:最⻓公共子序列是 "ace"，它的⻓度为 3。
//示例 2: 输入:text1 = "abc", text2 = "abc" 输出:3
//解释:最⻓公共子序列是 "abc"，它的⻓度为 3。
//示例 3: 输入:text1 = "abc", text2 = "def" 输出:0 解释:两个字符串没有公共子序列，返回 0。

func longestCommonSubsequence1(text1, text2 string) int {
	l1, l2 := len(text1), len(text2)
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}
	var result int
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			if dp[i][j] > result {
				result = dp[i][j]
			}
		}
	}
	return result
}
func longestCommonSubsequence(text1, text2 string) int {
	l1, l2 := len(text1), len(text2)
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}
	var result int
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
			if result < dp[i][j] {
				result = dp[i][j]
			}
		}
	}
	return result
}

func Handle44() {
	text1 := "abc"
	text2 := "def"
	fmt.Println(longestCommonSubsequence(text1, text2))
}
