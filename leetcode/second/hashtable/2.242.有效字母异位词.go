package hashtable

import "fmt"

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//示例 1: 输入: s = "anagram", t = "nagaram" 输出: true
//示例 2: 输入: s = "rat", t = "car" 输出: false
//说明: 你可以假设字符串只包含小写字母。
//若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

func isAnagram(s string, t string) bool {
	var record [26]int

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			record[s[i]-'a']++
			record[t[i]-'a']--
		}
	}
	for _, v := range record {
		if v != 0 {
			return false
		}
	}

	return true
}

func Handle2() {
	s := "rat"
	t := "car"
	res := isAnagram(s, t)
	fmt.Println(res)
}
