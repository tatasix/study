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
	var count int
	result := math.MinInt32
	l := len(nums)
	for i := 0; i < l; i++ {
		count += nums[i]
		if result < count {
			result = count
		}
		if count < 0 {
			count = 0
		}
	}
	return result
}

func maxSubArray(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	dp[0] = nums[0]
	result := nums[0]
	for i := 1; i < l; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}

func Handle46() {
	nums := []int{1}
	fmt.Println(maxSubArray(nums))
}
