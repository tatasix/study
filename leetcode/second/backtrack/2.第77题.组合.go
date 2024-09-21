package backtrack

import "fmt"

//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//
//示例: 输入: n = 4, k = 2 输出: [ [2,4], [3,4], [2,3], [1,2], [1,3], [1,4], ]

func combine(n int, k int) [][]int {
	var res [][]int
	var path []int
	var backTrack func(startIndex int)
	backTrack = func(startIndex int) {
		l := len(path)
		if l == k {
			temp := make([]int, l)
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := startIndex; i <= n; i++ {
			path = append(path, i)
			backTrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTrack(1)
	return res
}

func Handle2() {
	fmt.Println(combine(4, 2))
}
