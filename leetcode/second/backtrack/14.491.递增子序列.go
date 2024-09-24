package backtrack

import (
	"fmt"
)

// 491.递增子序列
// 给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的⻓度至少是2。 示例:
// 输入: [4, 6, 7, 7]
// 输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
func findSubsequences(nums []int) [][]int {
	var res [][]int
	var path []int
	l1 := len(nums)
	var backTrack func(start int)
	backTrack = func(start int) {
		if l1 < start {
			return
		}
		l2 := len(path)
		if l2 > 1 {
			temp := make([]int, l2)
			copy(temp, path)
			res = append(res, temp)
		}
		used := make([]bool, 201)
		for i := start; i < l1; i++ {
			current := nums[i]
			l3 := len(path)
			if used[current+100] || (l3 > 0 && path[l3-1] > current) {
				continue
			}
			used[current+100] = true
			path = append(path, current)
			backTrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return res
}

func Handle14() {
	fmt.Println(findSubsequences([]int{-100, -99, -98, -97, -96, -96}))
}

//[4,4,3,2,1]
//输出
//[[4,4],[4,4,3],[4,4,3,2],[4,4,3,2,1],[4,4,3,1],[4,4,2],[4,4,2,1],[4,4,1],[4,3],[4,3,2],[4,3,2,1],[4,3,1],[4,2],[4,2,1],[4,1],[3,2],[3,2,1],[3,1],[2,1]]
//预期结果
//[[4,4]]
//nums = [-100,-99,-98,-97,-96,-96]
