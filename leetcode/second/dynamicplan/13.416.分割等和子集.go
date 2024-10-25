package dynamicplan

import "fmt"

// 给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
// 注意: 每个数组中的元素不会超过 100 数组的大小不会超过 200
// 示例 1:输入: [1, 5, 11, 5]输出: true 解释: 数组可以分割成 [1, 5, 5] 和 [11].
// 示例 2:输入: [1, 2, 3, 5]输出: false 解释: 数组不能分割成两个元素和相等的子集.
func canPartition1(nums []int) bool {
	var sum int
	l := len(nums)
	for i := 0; i < l; i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	dp[0] = 0
	dp[1] = 0
	for i := 0; i < l; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}

	return dp[target] == target
}
func canPartition(nums []int) bool {
	var sum int
	l := len(nums)
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	dp[0] = 0
	dp[1] = 0
	for i := 0; i < l; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}

func Handle13() {
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
}
