package string

import "fmt"

// 给定一个字符串，逐个翻转字符串中的每个单词。
//
// 示例 1：
// 输入: "the sky is blue"
// 输出: "blue is sky the"
//
// 示例 2：
// 输入: "  hello world!  "
// 输出: "world! hello"
// 解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
//
// 示例 3：
// 输入: " a good   example  "
// 输出: "example good a"
// 解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
func reverseWords(s string) string {
	l := len(s)
	var res string
	right := l - 1
	for left := l - 1; left >= 0; left-- {
		if s[right] == ' ' {
			right--
			continue
		}
		if s[left] == ' ' && left != right {
			res += string(s[left+1:right+1]) + " "
			if left > 0 {
				right = left - 1
			} else {
				right = left
			}
		}
		if left == 0 && s[left] != ' ' {
			res += s[left : right+1]
		}
	}
	newL := len(res)
	if res[newL-1] == ' ' {
		res = res[:newL-1]
	}
	return res
}
func Handle4() {
	res := reverseWords("the sky is blue")
	fmt.Println(res)
}
