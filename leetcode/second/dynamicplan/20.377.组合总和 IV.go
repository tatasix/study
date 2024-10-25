package dynamicplan

import "fmt"

//给定一个由正整数组成且不存在重复数字的数组，找出和为给定目标正整数的组合的个数。 示例:
//nums = [1, 2, 3] target = 4
//所有可能的组合为:
//(1, 1, 1, 1) (1, 1, 2) (1, 2, 1) (1, 3)
//(2, 1, 1) (2, 2) (3, 1)
//请注意，顺序不同的序列被视作不同的组合。 因此输出为 7。

func combinationSum5(nums []int, target int) int {
	l := len(nums)
	dp := make([]int, target+1)
	dp[0] = 1

	for j := 0; j <= target; j++ {
		for i := 0; i < l; i++ {
			if j < nums[i] {
				continue
			}
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[target]
}
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j < nums[i] {
				continue
			}
			dp[j] += dp[j-nums[i]]
		}
		fmt.Println(dp)
	}
	return dp[target]
}
func Handle20() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}
