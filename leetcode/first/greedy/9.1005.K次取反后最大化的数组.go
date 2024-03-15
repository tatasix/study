package greedy

import (
	"fmt"
	"math"
	"sort"
)

func largestSumAfterKNegations1(nums []int, k int) int {
	var result int
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})
	l := len(nums)
	for i := 0; i < l; i++ {
		if k > 0 && nums[i] < 0 {
			nums[i] *= -1
			k--
		}
	}
	if k > 0 && k%2 == 1 {
		nums[l-1] *= -1
	}
	for i := 0; i < l; i++ {
		result += nums[i]
	}
	return result
}

func largestSumAfterKNegations(nums []int, k int) int {
	var result int
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] < 0 && k > 0 {
			result -= nums[i]
			k--
		} else {
			result += nums[i]
		}
	}
	if k > 0 && k%2 == 1 {
		last := math.Abs(float64(nums[l-1]))
		result -= int(last + last)
	}

	return result
}

//示例 1：
//
//输入：A = [4,2,3], K = 1
//输出：5
//解释：选择索引 (1,) ，然后 A 变为 [4,-2,3]。
//示例 2：
//
//输入：A = [3,-1,0,2], K = 3
//输出：6
//解释：选择索引 (1, 2, 2) ，然后 A 变为 [3,1,0,2]。
//示例 3：
//
//输入：A = [2,-3,-1,5,-4], K = 2
//输出：13
//解释：选择索引 (1, 4) ，然后 A 变为 [2,3,-1,5,4]。

func Handle9() {
	nums := []int{2, -3, -1, 5, -4}
	fmt.Println(largestSumAfterKNegations(nums, 2))
}
