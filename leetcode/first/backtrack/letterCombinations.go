package backtrack

func letterCombinations(digits string) []string {
	m := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	var result []string
	var path []byte
	if digits == "" {
		return result
	}

	length := len(digits)
	var backtrack func(start int)
	backtrack = func(start int) {
		if length == len(path) {
			tmp := string(path)
			result = append(result, tmp)
			return
		}
		next := int(digits[start] - '0')
		str := m[next-2]
		for i := 0; i <= len(str); i++ {
			path = append(path, str[i])
			backtrack(start + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return result
}
