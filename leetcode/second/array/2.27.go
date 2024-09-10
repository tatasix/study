package array

import "fmt"

func removeElement(nums []int, val int) int {
	left, right := 0, 0
	for ; right < len(nums); right++ {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}

	return left
}

func Handle3() {
	nums := []int{-1, 6, 1, 4, 20, 6, 9}
	target := 6
	removeElement(nums, target)
	fmt.Println(nums)
}
