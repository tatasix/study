package dynamicplan

import "fmt"

// 给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个
// 在字典中出现的单词。
// 说明:
// 拆分时可以重复使用字典中的单词。
// 你可以假设字典中没有重复的单词。
// 示例 1: 输入: s = "leetcode", wordDict = ["leet", "code"] 输出: true
// 解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
// 示例 2: 输入: s = "applepenapple", wordDict = ["apple", "pen"] 输出: true
// 解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
// 注意你可以重复使用字典中的单词。
// 示例 3: 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"] 输出: false
func wordBreak1(s string, wordDict []string) bool {
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
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]struct{}, len(wordDict))
	for _, v := range wordDict {
		wordMap[v] = struct{}{}
	}
	l := len(s)
	dp := make([]bool, l+1)
	dp[0] = true
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			_, ok := wordMap[s[j:i]]
			if dp[j] && ok {
				dp[i] = true
			}
		}
	}
	return dp[l]
}
func Handle26() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}
