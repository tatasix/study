package array

import "fmt"

// int minSubArrayLen(int s, vector<int>& nums) {
// 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s
// 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
// 示例：
// 输入：s = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
func minSubArrayLen(s int, nums []int) int {
	sum := 0
	l := len(nums)
	left, right := 0, 0
	minLength := l + 1
	for ; right < l; right++ {
		sum += nums[right]
		for sum >= s {
			if minLength > right-left+1 {
				minLength = right - left + 1
			}
			sum -= nums[left]
			left++
		}

	}
	if minLength == l+1 {
		return 0
	}
	return minLength
}

func Handle5() {
	s := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(s, nums))
}
func minSubArrayLen1(target int, nums []int) int {
	left, right, sum, min := 0, 0, 0, len(nums)+1

	for ; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			subLength := right - left + 1
			if subLength < min {
				min = subLength
			}
			sum -= nums[left]
			left++
		}
	}
	if min == len(nums)+1 {
		return 0
	}
	return min
}
