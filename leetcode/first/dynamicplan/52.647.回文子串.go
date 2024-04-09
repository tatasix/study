package dynamicplan

import "fmt"

// 给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
func countSubstrings(s string) int {
	var result int
	l := len(s)
	dp := make([][]bool, l)
	for i := range dp {
		dp[i] = make([]bool, l)
	}
	for i := l - 1; i >= 0; i-- {
		for j := i; j < l; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					result++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					result++
					dp[i][j] = true
				}
			}

		}
	}
	return result
}
func Handle52() {
	s := "aaa"
	fmt.Println(countSubstrings(s))
}
