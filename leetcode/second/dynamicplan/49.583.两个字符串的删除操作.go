package dynamicplan

import (
	"fmt"
)

// 给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，
// 每步可以删除任意一个字符串中的一个字符。
// 示例： 输入: "sea", "eat" 输出: 2
// 解释: 第一步将"sea"变为"ea"，第二步将"eat"变为"ea"
func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}
	dp[0][0] = 0
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1]+2, min(dp[i-1][j], dp[i][j-1])+1)
			}
		}
	}

	return dp[l1][l2]
}

func Handle49() {
	s1, s2 := "a", "b"
	fmt.Println(minDistance(s1, s2))
}
