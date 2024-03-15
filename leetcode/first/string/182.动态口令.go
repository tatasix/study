package string

func DynamicPassword(password string, target int) string {
	ss := []byte(password)
	ss = append(ss, ss[:target]...)
	return string(ss[target:])
}
