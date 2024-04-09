package dynamicplan

import "fmt"

// 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
// 你可以对一个单词进行如下三种操作：
// 插入一个字符
// 删除一个字符
// 替换一个字符
// 示例 1： 输入：word1 = "horse", word2 = "ros" 输出：3
// 解释： horse -> rorse (将 'h' 替换为 'r') rorse -> rose (删除 'r') rose -> ros (删除 'e')
// 示例 2： 输入：word1 = "intention", word2 = "execution" 输出：5
func minDistance1(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
		dp[i][0] = i
	}
	for j := 1; j <= l2; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
	return dp[l1][l2]
}

func Handle50() {
	word1, word2 := "intention", "execution"
	fmt.Println(minDistance1(word1, word2))
}
