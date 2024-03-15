package string

func reverseString(s []string) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		right--
		left++
	}

}
