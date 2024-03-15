package greedy

import "fmt"

func wiggleMaxLength(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}
	count, pre := 1, 0

	for i := 0; i < l-1; i++ {
		next := nums[i+1] - nums[i]
		if (pre <= 0 && next > 0) || (pre >= 0 && next < 0) {
			count++
			pre = next
		}

	}

	return count
}

func Handle3() {
	fmt.Println(wiggleMaxLength([]int{0, 0}))
}

//输入：nums = [1,7,4,9,2,5]
//输出：6
//解释：整个序列均为摆动序列，各元素之间的差值为 (6, -3, 5, -7, 3) 。
//示例 2：
//
//输入：nums = [1,17,5,10,13,15,10,5,16,8]
//输出：7
//解释：这个序列包含几个长度为 7 摆动序列。
//其中一个是 [1, 17, 10, 13, 10, 16, 8] ，各元素之间的差值为 (16, -7, 3, -3, 6, -8) 。
//示例 3：
//
//输入：nums = [1,2,3,4,5,6,7,8,9]
//输出：2
