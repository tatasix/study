package backtrack

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	var sum int
	sort.Ints(candidates)
	l := len(candidates)
	used := make([]bool, l)
	var backTrack func(start int)
	backTrack = func(start int) {
		if sum > target {
			return
		}
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := start; i < l; i++ {
			if i > 0 && candidates[i] == candidates[i-1] && !used[i-1] {
				continue
			}
			sum += candidates[i]
			used[i] = true
			path = append(path, candidates[i])
			backTrack(i + 1)
			path = path[:len(path)-1]
			used[i] = false
			sum -= candidates[i]
		}
	}
	backTrack(0)
	return res
}
func Handle8() {
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
