package dynamicplan

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
		fmt.Println(dp)
	}

	return dp[n]
}

func Handle24() {
	fmt.Println(numSquares(25))
}
