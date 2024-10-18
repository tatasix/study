package dynamicplan

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	l := len(s)
	l1 := len(wordDict)
	wordDictMap := make(map[string]bool, l1)
	for i := 0; i < l1; i++ {
		wordDictMap[wordDict[i]] = true
	}
	dp := make([]bool, l+1)
	dp[0] = true

	for j := 1; j <= l; j++ {
		for i := 0; i < j; i++ {
			if dp[i] && wordDictMap[s[i:j]] {
				dp[j] = true
			}
		}
	}
	return dp[l]
}

func Handle26() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}
