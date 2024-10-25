package dynamicplan

import "fmt"

// 给两个整数数组 A 和 B ，返回两个数组中公共的、⻓度最⻓的子数组的⻓度。
// 示例: 输入: A: [1,2,3,2,1] B: [3,2,1,4,7] 输出:3
// 解释: ⻓度最⻓的公共子数组是 [3, 2, 1] 。
func findLength1(nums1 []int, nums2 []int) int {
	l1, l2 := len(nums1), len(nums2)
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}
	var result int
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > result {
					result = dp[i][j]
				}
			}
		}
	}
	return result
}

func findLength(nums1 []int, nums2 []int) int {
	l1, l2 := len(nums1), len(nums2)
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}
	var result int
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > result {
					result = dp[i][j]
				}
			}
		}
	}
	return result
}

func Handle43() {
	A := []int{1, 2, 3, 2, 1}
	B := []int{3, 2, 1, 4, 7}
	fmt.Println(findLength(A, B))
}
