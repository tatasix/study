package hash

import (
	"fmt"
	"sort"
)

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

func Handle() {
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))

}
