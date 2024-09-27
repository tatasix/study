package greedy

import "fmt"

//55. 跳跃游戏
//给定一个非负整数数组，你最初位于数组的第一个位置。数组中的每个元素代表你在该位置可以跳跃的最大长度。
//判断你是否能够到达最后一个位置。
//示例 1: 输入: [2,3,1,1,4] 输出: true
//示例 2: 输入: [3,2,1,0,4] 输出: false

func canJump(nums []int) bool {
	var cover int
	l := len(nums)
	for i := 0; i < cover; i++ {
		if cover >= l-1 {
			return true
		}
		if i+nums[i] > cover {
			cover = i + nums[i]
		}
	}
	return false
}

func Handle7() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
