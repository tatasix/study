package backtrack

import "fmt"

// 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
// 输入: [1,2,3] 输出: [ [1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1] ]
func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	l := len(nums)
	used := make([]bool, l)
	var backTrack func(start int)
	backTrack = func(start int) {
		if len(path) == l {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := 0; i < l; i++ {
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			backTrack(0)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return res
}

func Handle15() {
	fmt.Println(permute([]int{1, 2, 3}))
}
