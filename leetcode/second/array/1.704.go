package array

import "fmt"

func search(nums []int, target int) int {
	l := len(nums)
	left, right := 0, l-1
	for left <= right {
		center := (left + right) >> 1
		if nums[center] < target {
			left = center + 1
		} else if nums[center] > target {
			right = center - 1
		} else {
			return center
		}
	}
	return -1
}

func Handle1() {
	nums := []int{-1, 1, 4, 6, 9}
	target := 10
	fmt.Println(search(nums, target))
}
