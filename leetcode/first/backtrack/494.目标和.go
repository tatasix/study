package backtrack

import "fmt"

// 给你一个非负整数数组 nums 和一个整数 target 。
//
// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//
// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
func findTargetSumWays(nums []int, target int) int {
	var (
		result    int
		sum       int
		backtrack func(start int)
	)
	backtrack = func(start int) {

		if start == len(nums) {
			if sum == target {
				result++
			}
			return
		}
		sum += nums[start]
		backtrack(start + 1)
		sum -= nums[start]
		sum -= nums[start]
		backtrack(start + 1)
		sum += nums[start]
	}
	backtrack(0)
	return result
}

func Handle() {
	nums := []int{1, 1, 1, 1, 1}
	target := 2
	fmt.Println(findTargetSumWays(nums, target))
}
