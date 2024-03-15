package greedy

import "strconv"

func monotoneIncreasingDigits(n int) int {
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
