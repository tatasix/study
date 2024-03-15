package string

import "fmt"

// 给定一个字符串 s 和一个整数 k，从字符串开头算起, 每计数至 2k 个字符，就反转这 2k 个字符中的前 k 个字符。
// 如果剩余字符少于 k 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
// 示例:
// 输入: s = "abcdefg", k = 2
// 输出: "bacdfeg"
// "aa" 6
func reverseStr(s string, k int) string {
	sByte := []byte(s)
	length := len(sByte)
	for i := 0; i < length; i = i + 2*k {
		if length < i+k-1 {
			reverse(sByte, i, length)
		} else {
			reverse(sByte, i, i+k-1)
		}
	}
	return string(sByte)
}

func reverse(s []byte, start, end int) {
	for l, r := start, end; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
	return
}

func Handle() {
	fmt.Println(reverseStr("abcdefg", 2))
}
