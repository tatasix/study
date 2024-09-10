package hashtable

import (
	"fmt"
	"sort"
)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
//使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//注意： 答案中不可以包含重复的三元组。
//示例：
//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为： [ [-1, 0, 1], [-1, -1, 2] ]

func threeSum2(nums []int) [][]int {
	var res [][]int
	return res
}
func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	length := len(nums)

	for i := 0; i < length; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l := i + 1
		r := length - 1
		for l < r {
			n2 := nums[l]
			n3 := nums[r]
			sum := n1 + n2 + n3
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				result = append(result, []int{n1, n2, n3})

				for l < r && nums[l] == n2 {
					l++
				}

				for l < r && nums[r] == n3 {
					r--
				}
			}

		}
	}
	return result
}

func Handle8() {
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(res)
}
