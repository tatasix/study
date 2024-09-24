package backtrack

import (
	"fmt"
	"sort"
)

// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集 幂集）。
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
// nums = [1,2,2]
// [[],[1],[1,2],[1,2,2],[1,2],[2],[2,2],[2]]
// 预期结果 [[],[1],[1,2],[1,2,2],[2],[2,2]]
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	var path []int
	sort.Ints(nums)
	l := len(nums)
	used := make(map[int]bool, l)
	var backTrack func(start int)
	backTrack = func(start int) {
		if start > l {
			return
		}
		l1 := len(path)
		temp := make([]int, l1)
		copy(temp, path)
		res = append(res, temp)
		for i := start; i < l; i++ {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backTrack(i + 1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backTrack(0)

	return res
}

func Handle13() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
