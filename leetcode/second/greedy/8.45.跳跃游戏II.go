package greedy

import "fmt"

//给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
//每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
//0 <= j <= nums[i]
//i + j < n
//返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。
//示例 1: 输入: nums = [2,3,1,1,4] 输出: 2
//示例 2: 输入: nums = [2,3,0,1,4] 输出: 2

func jump(nums []int) int {
	var cover int
	var res int
	var next int
	l := len(nums)
	if l == 1 {
		return 0
	}
	for i := 0; i < l; i++ {
		current := i + nums[i]
		next = max(next, current)
		if i == cover {
			res++
			cover = next
		}
		if cover >= l-1 {
			break
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Handle8() {
	fmt.Println(jump([]int{1}))
}
func jump1(nums []int) int {
	var (
		result  int
		current int
		next    int
	)
	l := len(nums)
	if l < 2 {
		return result
	}

	for i := 0; i < l-1; i++ {
		temp := i + nums[i]
		if temp > next {
			next = temp
		}
		if i == current {
			current = next
			result++
		}

	}
	return result

}
