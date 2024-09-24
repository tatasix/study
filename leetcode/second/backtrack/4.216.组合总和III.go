package backtrack

import "fmt"

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var path []int
	var sum int
	var backTrack func(start int)
	backTrack = func(start int) {
		if len(path) >= k || sum >= n {
			if sum == n && len(path) == k {
				temp := make([]int, k)
				copy(temp, path)
				res = append(res, temp)
			}
			return
		}
		for i := start; i < 10; i++ {
			path = append(path, i)
			sum += i
			backTrack(i + 1)
			path = path[:len(path)-1]
			sum -= i
		}
	}
	backTrack(1)
	return res
}

func Handle4() {
	fmt.Println(combinationSum3(3, 9))
}
