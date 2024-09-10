package hashtable

import (
	"fmt"
	"sort"
)

// 题意：给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在
// 四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
// 注意：
// 答案中不可以包含重复的四元组。
// 示例： 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。 满足要求的四元组集合为：
// [ [-1, 0, 0, 1], [-2, -1, 1, 2], [-2, 0, 0, 2] ]
// [-2, -1, 0,0, 1,  2]
func fourSum(nums []int, target int) [][]int {
	var res [][]int
	n := len(nums)
	if n < 4 {
		return res
	}
	sort.Ints(nums)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := n - 1
			for left < right {
				nl := nums[left]
				nr := nums[right]
				sum := nums[i] + nums[j] + nl + nr
				if sum > target {
					right--
				} else if sum < target {
					left++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[right] == nr {
						right--
					}
					for left < right && nums[left] == nl {
						left++
					}
				}
			}
		}
	}
	return res
}

func fourSum1(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length-3; i++ {
		n1 := nums[i]
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		for k := i + 1; k < length-2; k++ {
			n2 := nums[k]
			if k > i+1 && n2 == nums[k-1] {
				continue
			}
			l := k + 1
			r := length - 1
			for l < r {
				n3, n4 := nums[l], nums[r]
				sum := n1 + n2 + n3 + n4
				if sum == target {
					res = append(res, []int{n1, n2, n3, n4})
					for l < r && nums[l] == n3 {
						l++
					}
					for l < r && nums[r] == n4 {
						r--
					}
				} else if sum > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return res
}

func Handle9() {
	//res := fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0)
	res := fourSum([]int{2, 2, 2, 2, 2}, 8)
	fmt.Println(res)
}
