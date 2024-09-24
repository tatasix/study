package backtrack

import "fmt"

// 分割回文串.
// 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
// 返回 s 所有可能的分割方案。
// 示例:
// 输入: "aab" 输出:
// [
// ["aa","b"],
// ["a","a","b"] ]
func isPartition(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// "aab"
// [  ["aa","b"],  ["a","a","b"] ]

func partition(s string) [][]string {
	var res [][]string
	var path []string
	l := len(s)
	var backTrack func(start int)
	backTrack = func(start int) {
		if start > l+1 {
			return
		}
		if start == l {
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := start; i < l; i++ {
			current := s[start : i+1]
			if !isPartition(current) {
				continue
			}
			path = append(path, current)
			backTrack(i + 1)
			path = path[:len(path)-1]
		}

	}
	backTrack(0)
	return res
}

func Handle9() {
	fmt.Println(partition("efe"))
}
