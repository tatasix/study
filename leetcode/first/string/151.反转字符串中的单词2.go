package string

func ReverseWords2(s string) string {
	ss := []byte(s)
	length := len(ss)
	left, right := length-1, length
	newS := make([]byte, length)
	for i := length - 1; i >= 0; i-- {
		if ss[i] == ' ' {
			newS = append(newS, ss[left:right]...)
			right = left
		}
		left--
	}
	newS = append(newS, ' ')
	newS = append(newS, ss[:right]...)
	return string(newS)
}
