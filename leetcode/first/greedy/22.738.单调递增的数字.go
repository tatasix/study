package greedy

import (
	"fmt"
	"strconv"
)

func monotoneIncreasingDigits1(n int) int {
	s := strconv.Itoa(n)
	ss := []byte(s)
	l := len(ss)
	if l < 2 {
		return n
	}
	start := l
	for i := l - 1; i > 0; i-- {
		if ss[i-1] > ss[i] {
			ss[i-1] -= 1
			start = i
		}
	}

	for i := start; i < l; i++ {
		ss[i] = '9'
	}
	res, _ := strconv.Atoi(string(ss))
	return res
}

// 给定一个非负整数 N，找出小于或等于 N 的最大的整数，
// 同时这个整数需要满足其各个位数上的数字是单调递增。
// （当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。）
// 示例 1:   输入: N = 10    输出: 9
// 示例 2:   输入: N = 1234  输出: 1234
// 示例 3:   输入: N = 332   输出: 299
func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	ss := []byte(s)
	start := len(ss)
	for i := len(ss) - 1; i > 0; i-- {
		if ss[i] < ss[i-1] {
			ss[i-1] -= 1
			start = i
		}
	}
	for i := start; i < len(ss); i++ {
		ss[i] = '9'
	}
	ssss, _ := strconv.Atoi(string(ss))
	return ssss
}
func Handle22() {
	fmt.Println(monotoneIncreasingDigits(100))
}
