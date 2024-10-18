package dynamicplan

import "fmt"

// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素
// 就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你不触动警报装置的情况下，一夜之内能够偷窃到的最高金额。
// 示例 1: 输入:[1,2,3,1] 输出:4
// 示例 2: 输入:[2,7,9,3,1] 输出:12
func rob1(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < l; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[l-1]
}

func Handle29() {
	fmt.Println(rob1([]int{2, 7, 9, 3, 1}))
}
