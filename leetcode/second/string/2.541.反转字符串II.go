package string

import "fmt"

// 给定一个字符串 s 和一个整数 k，从字符串开头算起, 每计数至 2k 个字符，就反转这 2k 个字符中的前 k 个字符。
// 如果剩余字符少于 k 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
// 示例:
// 输入: s = "abcdefg", k = 2
// 输出: "bacdfeg"
func reverseStr(s string, k int) string {
	s1 := []byte(s)
	l := len(s)
	for i := 0; i < l; i += 2 * k {
		right := k
		if right > l {
			right = l
		}
		reverseString(s1[i : i+right])
	}
	return string(s1)
}

func Handle2() {
	res := reverseStr("a", 2)
	fmt.Println(res)
}
