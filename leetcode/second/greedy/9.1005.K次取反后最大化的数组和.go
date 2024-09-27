package greedy

import (
	"fmt"
	"math"
	"sort"
)

//给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：
//选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
//重复这个过程恰好 k 次。可以多次选择同一个下标 i 。
//以这种方式修改数组后，返回数组 可能的最大和 。
//示例 1：输入：nums = [4,2,3], k = 1 输出：5
//解释：选择下标 1 ，nums 变为 [4,-2,3] 。
//示例 2： 输入：nums = [3,-1,0,2], k = 3 输出：6
//解释：选择下标 (1, 2, 2) ，nums 变为 [3,1,0,2] 。
//示例 3：输入：nums = [2,-3,-1,5,-4], k = 2 输出：13
//解释：选择下标 (1, 4) ，nums 变为 [2,3,-1,5,4] 。

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	var res int
	l := len(nums)
	index := getMiddle(nums)
	if index == 0 {
		for i := 0; i < l; i++ {
			res += nums[i]
		}
		if k%2 == 1 {
			res -= nums[0] * 2
		}
		return res
	}
	if index == -1 {
		if k >= l {
			for i := 0; i < l; i++ {
				res -= nums[i]
			}
			if (k-l)%2 == 1 {
				res += nums[l-1] * 2
			}
		} else {
			j := k
			for i := 0; i < l; i++ {
				if j == 0 {
					res += nums[i]
				} else {
					res -= nums[i]
					j--
				}
			}
		}

		return res
	}
	if k > index {
		for i := 0; i < index; i++ {
			res -= nums[i]
		}
		for i := index; i < l; i++ {
			res += nums[i]
		}

		if (k-index)%2 == 1 {
			if -nums[index-1] > nums[index] {
				res -= nums[index] * 2
			} else {
				res += nums[index-1] * 2
			}
		}
	} else {
		j := k
		for i := 0; i < l; i++ {
			if j == 0 {
				res += nums[i]
			} else {
				res -= nums[i]
				j--
			}
		}
	}

	return res
}

func getMiddle(nums []int) int {
	if nums[0] >= 0 {
		return 0
	}
	if nums[len(nums)-1] < 0 {
		return -1
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < 0 && nums[i+1] >= 0 {
			return i + 1
		}
	}
	return -1
}

func Handle9() {
	fmt.Println(largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2))
}

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
