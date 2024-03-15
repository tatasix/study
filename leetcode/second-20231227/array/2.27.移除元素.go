package array

import "fmt"

func remove(nums []int, val int) int {

	var left, right int

	for right = 0; right < len(nums); right++ {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}

func Handle2() {
	res := remove([]int{3, 2, 2, 3}, 3)
	fmt.Println(res)
}
