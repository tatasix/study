package dynamicplan

import (
	"fmt"
)

// 给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
// 注意: 每个数组中的元素不会超过 100 数组的大小不会超过 200
//
// 示例 1:
// 输入: [1, 5, 11, 5]
// 输出: true
// 解释: 数组可以分割成 [1, 5, 5] 和 [11].
// 示例 2:
// 输入: [1, 2, 3, 5]
// 输出: false
// 解释: 数组不能分割成两个元素和相等的子集.
func canPartition1(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	for _, num := range nums {
		for j := target; j >= num; j-- {
			if dp[j] < dp[j-num]+num {
				dp[j] = dp[j-num] + num
			}
		}
		fmt.Println(dp)
	}
	return dp[target] == target
}

func Handle13() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}
func canPartition(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	for _, num := range nums {
		for j := target; j >= num; j-- {
			if dp[j] > dp[j-num]+num {
				dp[j] = dp[j-num] + num
			}
		}
	}
	return dp[target] == target
}
