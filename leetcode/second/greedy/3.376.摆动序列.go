package greedy

import "fmt"

// 给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列的长度 。
// 示例 1： 输入：nums = [1,7,4,9,2,5] 输出：6
// 示例 2： 输入：nums = [1,17,5,10,13,15,10,5,16,8] 输出：7
// 示例 3： 输入：nums = [1,2,3,4,5,6,7,8,9] 输出：2

func wiggleMaxLength(nums []int) int {
	res := 1
	l := len(nums)
	if l < 2 {
		return l
	}
	preDiff := 0
	for i := 0; i < l-1; i++ {
		curDiff := nums[i+1] - nums[i]
		if (preDiff <= 0 && curDiff > 0) || (preDiff >= 0 && curDiff < 0) {
			res++
			preDiff = curDiff
		}
	}
	return res
}

func Handle3() {
	fmt.Println(wiggleMaxLength([]int{0, 0}))
}
