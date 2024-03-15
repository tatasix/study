package dynamicplan

import (
	"fmt"
)

func canPartition(nums []int) bool {
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
