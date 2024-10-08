package string

import "fmt"

//给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
//
//示例 1: 输入: haystack = "hello", needle = "ll" 输出: 2
//
//示例 2: 输入: haystack = "aaaaa", needle = "bba" 输出: -1
//
//说明: 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。

func strStr1(haystack, needle string) int {
	if needle == "" {
		return 0
	}
	lh := len(haystack)
	for i := 0; i < lh; i++ {
		var flag bool
		if haystack[i] == needle[0] {
			for j := 1; j < len(needle); j++ {
				if i+j >= lh || needle[j] != haystack[i+j] {
					flag = true
					break
				}
			}
			if !flag {
				return i
			}
		}
	}
	return -1
}

func strStr(haystack, needle string) int {
	if needle == "" {
		return 0
	}
	lh, ln := len(haystack), len(needle)

	for i := 0; i < lh; i++ {
		if i+ln <= lh && haystack[i:i+ln] == needle {
			return i
		}
	}
	return -1
}

func Handle() {
	res := strStr("a", "a")
	fmt.Println(res)
}
