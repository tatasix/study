package dynamicplan

import "fmt"

func lastStoneWeightII1(stones []int) int {
	// 15001 = 30 * 1000 /2 +1

	// 求target
	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2
	fmt.Println(target)
	dp := make([]int, target+1)
	// 遍历顺序
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			// 推导公式
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
		fmt.Println(dp)
	}
	return sum - 2*dp[target]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Handle14() {
	nums := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeightII(nums))
}
func lastStoneWeightII(stones []int) int {
	var sum int
	for _, v := range stones {
		sum += v
	}
	target := sum / 2
	dp := make([]int, target+1)
	for _, stone := range stones {
		for j := target; j >= stone; j-- {
			dp[j] = max(dp[j], dp[j-stone]+stone)
		}
	}

	return sum - 2*dp[target]
}
