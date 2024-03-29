package backtrack

import "fmt"

// 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//
// 示例: 输入: n = 4, k = 2 输出: [ [2,4], [3,4], [2,3], [1,2], [1,3], [1,4], ]
func combine(n int, k int) [][]int {
	var (
		result    [][]int
		path      []int
		backtrack func(start int)
	)
	backtrack = func(start int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := start; i <= n; i++ {
			path = append(path, i)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(1)
	return result
}

func Handle() {
	fmt.Println(combine(4, 2))
}
