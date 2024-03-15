package string

func ReverseWords(s string) string {
	ss := []byte(s)
	length := len(ss)
	left := 0
	right := 0
	for i := 0; i < length; i++ {
		if ss[i] == ' ' {
			reverse(ss[left:right])
			right++
			left = right
		} else {
			right++
		}
	}
	return string(ss)
}

func reverse(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		right--
		left++
	}

}
