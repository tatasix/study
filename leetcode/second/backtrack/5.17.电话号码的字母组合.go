package backtrack

import "fmt"

func letterCombinations(digits string) []string {
	var res []string
	if digits == "" {
		return res
	}
	var path []byte
	numberMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var backTrack func(start int)
	backTrack = func(start int) {
		l := len(path)
		if l == len(digits) {
			temp := make([]byte, l)
			copy(temp, path)
			res = append(res, string(temp))
			return
		}
		str := numberMap[digits[start]]

		for _, v := range str {
			path = append(path, byte(v))
			backTrack(start + 1)
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return res
}

func Handle5() {
	fmt.Println(letterCombinations("23"))
}
