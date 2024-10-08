package dynamicplan

import (
	"fmt"
)

// 给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
// 说明：
// 拆分时可以重复使用字典中的单词。
// 你可以假设字典中没有重复的单词。
// 示例 1： 输入: s = "leetcode", wordDict = ["leet", "code"] 输出: true
// 解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
// 示例 2： 输入: s = "applepenapple", wordDict = ["apple", "pen"] 输出: true
// 解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
// 注意你可以重复使用字典中的单词。
// 示例 3： 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"] 输出: false
func wordBreak1(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, v := range wordDict {
		wordMap[v] = true
	}
	l := len(s)
	dp := make([]bool, l+1)
	dp[0] = true

	for i := 1; i <= l; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
			}
		}
		fmt.Println(dp)
	}
	return dp[l]
}
func Handle26() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict))
}
func wordBreak2(s string, wordDict []string) bool {

	wordMap := make(map[string]bool)
	for _, v := range wordDict {
		wordMap[v] = true
	}
	l := len(s)
	dp := make([]bool, l+1)
	dp[0] = true
	for i := 1; i <= l; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
			}
		}
	}
	return dp[l]
}
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, v := range wordDict {
		wordMap[v] = true
	}
	l := len(s)
	memory := make([]int, l)
	for i := range memory {
		memory[i] = -1
	}
	var backTrack func(startIndex int) bool
	backTrack = func(startIndex int) bool {
		if startIndex >= l {
			return true
		}
		if memory[startIndex] != -1 {
			return memory[startIndex] == 1
		}
		for i := startIndex; i < l; i++ {
			if wordMap[s[startIndex:i+1]] && backTrack(i+1) {
				memory[startIndex] = 1
				return true
			}
		}
		memory[startIndex] = 0
		return false
	}
	return backTrack(0)
}
