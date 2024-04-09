package dynamicplan

import (
	"fmt"
	"math"
)

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 子数组 是数组中的一个连续部分。
// 示例 1： 输入：nums = [-2,1,-3,4,-1,2,1,-5,4] 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
// 示例 2： 输入：nums = [1] 输出：1
func maxSubArray1(nums []int) int {
	count, result := 0, math.MinInt32
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > result {
			result = count
		}
		if count < 0 {
			count = 0
		}
	}

	return result
}

func maxSubArray(nums []int) int {
	result := nums[0]
	dp := make([]int, len(nums)+1)
	dp[0] = nums[0]
	for i := 1; i <= len(nums); i++ {
		dp[i] = math.MinInt32
	}
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}

func Handle46() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
