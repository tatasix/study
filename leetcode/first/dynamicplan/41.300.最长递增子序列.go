package dynamicplan

import "fmt"

func lengthOfLIS(nums []int) int {
	var result int
	l := len(nums)
	if l < 2 {
		return l
	}
	dp := make([]int, l)
	for i := range dp {
		dp[i] = 1
	}
	result = 1
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				if dp[i] > result {
					result = dp[i]
				}
			}
		}
	}
	return result
}

func Handle41() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}
