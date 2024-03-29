package dynamicplan

import "fmt"

// 给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。
// 返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
// 示例：
// 输入：nums: [1, 1, 1, 1, 1], S: 3
// left - right = s
// left + right = sum
// left = (s+sum)/2
// 输出：5
func findTargetSumWays1(nums []int, target int) int {
	var sum, t int
	for _, v := range nums {
		sum += v
	}
	if target > sum {
		return 0
	}
	if (target+sum)%2 == 1 {
		return 0
	}
	t = (sum + target) / 2
	dp := make([]int, t+1)
	dp[0] = 1
	for _, num := range nums {
		for j := t; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[t]
}

func Handle16() {
	nums := []int{1, 1, 1, 1, 1}
	s := 3
	fmt.Println(findTargetSumWays1(nums, s))
}
