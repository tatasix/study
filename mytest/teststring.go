package mytest

func TestString(a, b int) int {
	if a < b {
		a = b
	}
	return a
}
