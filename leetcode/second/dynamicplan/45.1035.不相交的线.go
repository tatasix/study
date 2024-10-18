package dynamicplan

import "fmt"

//在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。
//现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足：
// nums1[i] == nums2[j]
//且绘制的直线不与任何其他连线（非水平线）相交。
//请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。
//以这种方法绘制线条，并返回可以绘制的最大连线数。
//输入：nums1 = [1,4,2], nums2 = [1,2,4]  输出：2
//示例 2： 输入：nums1 = [2,5,1,2,5], nums2 = [10,5,2,1,5,2]  输出：3
//示例 3： 输入：nums1 = [1,3,7,1,7,5], nums2 = [1,9,2,5,1]  输出：2

func maxUncrossedLines(nums1 []int, nums2 []int) int {
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
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
			if result < dp[i][j] {
				result = dp[i][j]
			}
		}
	}
	return result
}

func Handle45() {
	nums1 := []int{1, 4, 2}
	nums2 := []int{1, 2, 4}
	fmt.Println(maxUncrossedLines(nums1, nums2))
}
