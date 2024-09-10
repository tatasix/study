package hashtable

import "fmt"

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那
// 两个 整数，并返回他们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
// 示例:
// 给定 nums = [2, 7, 11, 15], target = 9
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]
func twoSum(nums []int, target int) []int {
	set := make(map[int]int)
	for k, v := range nums {
		vv, ok := set[v]
		if ok {
			return []int{vv, k}
		}
		set[target-v] = k
	}
	return []int{}
}
func twoSum1(nums []int, target int) []int {
	l := len(nums)
	used := make(map[int]int, l)
	for i := 0; i < l; i++ {
		if v, ok := used[target-nums[i]]; ok {
			return []int{v, i}
		} else {
			used[nums[i]] = i
		}
	}
	return []int{}
}
func Handle5() {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)
}
