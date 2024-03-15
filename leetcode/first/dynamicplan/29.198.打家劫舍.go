package dynamicplan

func rob(nums []int) int {
	l := len(nums)
	dp := make([]int, l+1)
	dp[1] = nums[0]
	for i := 2; i <= l; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[l]
}
