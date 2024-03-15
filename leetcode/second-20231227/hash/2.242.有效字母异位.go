package hash

import "fmt"

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
//示例 1: 输入: s = "anagram", t = "nagaram" 输出: true
//示例 2: 输入: s = "rat", t = "car" 输出: false
//说明: 你可以假设字符串只包含小写字母。

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	record := [26]int{}
	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			continue
		}
		record[s[i]-'a']++
		record[t[i]-'a']--
	}
	for _, v := range record {
		if v != 0 {
			return false
		}
	}
	return true
}
func Handle2() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
