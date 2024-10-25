package dynamicplan

import "fmt"

// 给你一个整数数组 nums ，找到其中最⻓严格递增子序列的⻓度。
// 子序列是由数组派生而来的序列，删除(或不删除)数组中的元素而不改变其余元素的顺序。例如， [3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
// 示例 1: 输入:nums = [10,9,2,5,3,7,101,18] 输出:4
// 解释:最⻓递增子序列是 [2,3,7,101]，因此⻓度为 4 。
// 示例 2: 输入:nums = [0,1,0,3,2,3] 输出:4
// 示例 3: 输入:nums = [7,7,7,7,7,7,7] 输出:1
func lengthOfLIS1(nums []int) int {
	l := len(nums)
	result := 1
	dp := make([]int, l)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > result {
			result = dp[i]
		}
	}

	return result
}
func lengthOfLIS(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	for i := range dp {
		dp[i] = 1
	}
	var result int
	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				if result < dp[i] {
					result = dp[i]
				}
			}
		}
	}
	return result
}

func Handle41() {
	nums := []int{2, 3, 7, 101}
	fmt.Println(lengthOfLIS(nums))
}
