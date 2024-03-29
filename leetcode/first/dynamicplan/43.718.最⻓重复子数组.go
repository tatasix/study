package dynamicplan

import "fmt"

//给两个整数数组 A 和 B ，返回两个数组中公共的、⻓度最⻓的子数组的⻓度。
//示例:
//输入: A: [1,2,3,2,1] B: [3,2,1,4,7] 输出:3
//解释: ⻓度最⻓的公共子数组是 [3, 2, 1] 。

func findLength1(A, B []int) int {
	a := len(A)
	b := len(B)
	dp := make([][]int, a+1)
	for i := range dp {
		dp[i] = make([]int, b+1)
	}
	dp[0][0] = 0
	result := 0
	for i := 1; i <= a; i++ {
		for j := 1; j <= b; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > result {
				result = dp[i][j]
			}
		}
	}
	return result
}
func findLength(A, B []int) int {
	result := 0
	a := len(A)
	b := len(B)
	dp := make([]int, b+1)
	for i := 1; i <= a; i++ {
		for j := 1; j <= b; j++ {
			if A[i-1] == B[j-1] {
				dp[j] = dp[j-1] + 1
			}
			if dp[j] > result {
				result = dp[j]
			}
		}
		fmt.Println(dp)
	}
	return result
}

func Handle43() {
	A := []int{3, 2, 1}
	B := []int{3, 2, 1, 4, 7}
	fmt.Println(findLength(A, B))
}
