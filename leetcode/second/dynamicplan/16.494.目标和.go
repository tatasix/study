package dynamicplan

import "fmt"

// a+b = sum
// a-b = s a=(sum+s)/2
// 给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。
// 返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
// 示例： 输入：nums: [1, 1, 1, 1, 1], S: 3 输出：5
func findTargetSumWays1(nums []int, target int) int {
	var sum int
	l := len(nums)
	for i := 0; i < l; i++ {
		sum += nums[i]
	}
	if (target+sum)%2 == 1 {
		return 0
	}
	left := (target + sum) / 2
	dp := make([]int, left+1)
	dp[0] = 1
	for i := 0; i < l; i++ {
		for j := left; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[left]
}
func findTargetSumWays(nums []int, target int) int {
	var sum int
	l := len(nums)
	for _, v := range nums {
		sum += v
	}
	if (target+sum)%2 == 1 {
		return 0
	}
	t := (sum + target) / 2

	dp := make([]int, t+1)
	dp[0] = 1
	for i := 0; i < l; i++ {
		for j := t; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[t]
}
func Handle16() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
