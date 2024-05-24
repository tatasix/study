package slice

import "testing"

func Test_slice(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	t.Logf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}
