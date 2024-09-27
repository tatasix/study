package greedy

import (
	"fmt"
	"math"
)

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 子数组  是数组中的一个连续部分。
// 示例 1： 输入：nums = [-2,1,-3,4,-1,2,1,-5,4] 输出：6
// 示例 2： 输入：nums = [1] 输出：1
// 示例 3： 输入：nums = [5,4,-1,7,8] 输出：23

func maxSubArray(nums []int) int {
	res := math.MinInt32
	var count int
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > res {
			res = count
		}
		if count < 0 {
			count = 0
		}
	}

	return res
}

func Handle4() {
	fmt.Println(maxSubArray([]int{-1}))
}
