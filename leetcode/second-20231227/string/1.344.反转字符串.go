package string

import "fmt"

func reverseString(s []byte) {

	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
}

func Handle1() {
	s := []byte{'a', 'b', 'c', 'd'}
	reverseString(s)
	fmt.Println(string(s))
}
