package dynamicplan

import "fmt"

// 给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
func countSubstrings1(s string) int {
	l := len(s)
	var result int
	dp := make([][]bool, l)
	for i := range dp {
		dp[i] = make([]bool, l)
	}
	for i := l - 1; i >= 0; i-- {
		for j := i; j < l; j++ {
			if s[i] == s[j] {
				if j-i <= 1 || dp[i+1][j-1] {
					dp[i][j] = true
					result++
				}
			}
		}
	}
	return result
}
func countSubstrings2(s string) int {
	var result int
	l := len(s)

	f := func(i, j int) {
		for i >= 0 && j < l && s[i] == s[j] {
			result++
			i--
			j++
		}
	}
	for i := 0; i < l; i++ {
		f(i, i)
		f(i, i+1)
	}
	return result
}
func countSubstrings(s string) int {
	l := len(s)
	dp := make([][]bool, l)
	for i := range dp {
		dp[i] = make([]bool, l)
	}
	var result int
	for i := l - 1; i >= 0; i-- {
		for j := i; j < l; j++ {
			if s[i] == s[j] {
				if j-1 <= 1 || dp[i+1][j-1] == true {
					dp[i][j] = true
					result++
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

// 最长回文字串
func longestPalindrome(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}
	var m int
	var result string
	f := func(i, j int) {
		for i >= 0 && j < l && s[i] == s[j] {
			if j-i >= m {
				m = j - i
				result = s[i : j+1]
			}
			i--
			j++
		}
	}
	for i := 0; i < l; i++ {
		f(i, i)
		f(i, i+1)
	}
	return result
}
