package array

import "fmt"

func sortedSquares(nums []int) []int {
	l := len(nums)
	result := make([]int, l, l)
	left, right, index := 0, l-1, l-1
	for left <= right {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			result[index] = nums[left] * nums[left]
			left++
		} else {
			result[index] = nums[right] * nums[right]
			right--
		}
		index--
	}
	return result
}

func Handle4() {
	nums := []int{-4, -1, 0, 3, 10}

	fmt.Println(sortedSquares(nums))
}
