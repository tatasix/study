package array

import "fmt"

func find(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}
	var get func(left, right int) int
	get = func(left, right int) int {
		if left > right {
			return -1
		}
		middle := (right + left) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			return get(left, middle-1)
		} else {
			return get(middle+1, right)
		}

	}
	return get(0, len(nums)-1)
}

func Handle1() {
	res := find([]int{-1, 0, 3, 5, 9, 12}, 9)
	fmt.Println(res)
}
