package backtrack

import (
	"fmt"
	"sort"
)

//78.子集
//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集
//nums = [1,2,3] 输出:[[3], [1], [2], [1,2,3], [1,3], [2,3], [1,2], [] ]

func subsets(nums []int) [][]int {
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

func Handle11() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
