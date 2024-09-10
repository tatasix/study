package stackqueue

import (
	"fmt"
	"strconv"
	"strings"
)

// 根据 逆波兰表示法，求表达式的值。
// 有效的运算符包括 + ,  - ,  * ,  / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
// 说明：
// 整数除法只保留整数部分。 给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
// 示例 1： 输入: ["2", "1", "+", "3", "*"] 输出: 9
// 示例 2： 输入: ["4", "13", "5", "/", "+"] 输出: 6
// 示例 3： 输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
// 输出: 22

func evalRPN(tokens []string) int {
	var data []int
	for i := 0; i < len(tokens); i++ {
		tokens[i] = strings.TrimSpace(tokens[i])
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			l := len(data)
			e1 := data[l-2]
			e2 := data[l-1]
			var result int
			switch tokens[i] {
			case "+":
				result = e1 + e2
			case "-":
				result = e1 - e2
			case "*":
				result = e1 * e2
			case "/":
				result = e1 / e2
			}
			data = data[:l-2]
			data = append(data, result)
		} else {
			e, _ := strconv.Atoi(tokens[i])
			data = append(data, e)
		}
	}
	return data[0]
}

func Handle6() {
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
}
