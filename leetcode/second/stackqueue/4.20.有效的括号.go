package stackqueue

import "fmt"

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串。
// 示例 1: 输入: "()" 输出: true
// 示例 2: 输入: "()[]{}" 输出: true
// 示例 3: 输入: "(]" 输出: false
// 示例 4: 输入: "([)]" 输出: false
// 示例 5: 输入: "{[]}" 输出: true

func isValid(s string) bool {
	var data []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			data = append(data, s[i])
		} else {
			l := len(data)
			if l == 0 {
				return false
			}
			if !check(data[l-1], s[i]) {
				return false
			}
			data = data[:l-1]
		}

	}

	return len(data) == 0
}

func check(a, b byte) bool {
	switch a {
	case '[':
		return b == ']'
	case '(':
		return b == ')'
	default:
		return b == '}'
	}
}

func Handle4() {
	//a := "()" // 输出: true
	//a := "()[]{}" // 输出: true
	//a := "(]" // 输出: false
	//a := "([)]" //输出: false
	a := "{[]}" // 输出: true
	res := isValid(a)
	fmt.Println(res)
}
