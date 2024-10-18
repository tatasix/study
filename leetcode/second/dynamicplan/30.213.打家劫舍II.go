package dynamicplan

import "fmt"

func rob2(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	one := get(nums[1:])
	two := get(nums[:l-1])
	return max(one, two)
}

func get(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	dp := make([]int, l)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < l; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[l-1]
}

func Handle30() {
	fmt.Println(rob2([]int{1, 2, 3}))
}
