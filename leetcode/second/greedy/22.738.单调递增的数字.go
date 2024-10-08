package greedy

import (
	"fmt"
	"strconv"
)

//给定一个非负整数 N，找出小于或等于 N 的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。 (当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。)
//示例 1: 输入: N = 10 输出: 9
//示例 2: 输入: N = 1234 输出: 1234
//示例 3: 输入: N = 332 输出: 299

func monotoneIncreasingDigits(N int) int {
	s := strconv.Itoa(N)
	l := len(s)
	ss := []byte(s)
	flag := l
	for i := l - 1; i > 0; i-- {
		if ss[i] < ss[i-1] {
			ss[i-1]--
			flag = i
		}
	}
	for i := flag; i < l; i++ {
		ss[i] = '9'
	}
	re, _ := strconv.Atoi(string(ss))
	return re
}

func Handle22() {
	fmt.Println(monotoneIncreasingDigits(332))
}
