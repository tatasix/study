package dynamicplan

import "fmt"

func findLengthOfLCIS1(nums []int) int {
	result := 1
	l := len(nums)
	if l < 2 {
		return l
	}
	dp := make([]int, l)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < l; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
			if dp[i] > result {
				result = dp[i]
			}
		}
	}
	return result
}

func Handle42() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS(nums))
}

// 贪心
func findLengthOfLCIS(nums []int) int {
	count, result := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			count++
		} else {
			count = 1
		}
		if count > result {
			result = count
		}
	}
	return result
}
