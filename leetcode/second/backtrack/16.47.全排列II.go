package backtrack

import (
	"fmt"
	"sort"
)

// 输入：nums = [1,1,2]
// 输出：  [[1,1,2], [1,2,1], [2,1,1]]
// 输入：nums = [1,2,3] 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
func permuteUnique(nums []int) [][]int {
	var res [][]int
	var path []int
	sort.Ints(nums)
	l := len(nums)
	used := make([]bool, l)
	var backTrack func()
	backTrack = func() {
		l2 := len(path)
		if l2 == l {
			temp := make([]int, l2)
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := 0; i < l; i++ {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if !used[i] {
				path = append(path, nums[i])
				used[i] = true
				backTrack()
				used[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	backTrack()
	return res
}

func Handle16() {
	fmt.Println(permuteUnique([]int{3, 3, 0, 3}))
}
