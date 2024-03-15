package tree

import "fmt"

type TreeNode6 struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMax(nums []int) (index int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[index] {
			index = i
		}
	}
	return
}

func constructMaximumBinaryTree(nums []int) *TreeNode {

	var traversal func(left, right int) *TreeNode

	traversal = func(left, right int) *TreeNode {
		if len(nums) <= 0 {
			return nil
		}
		if left <= right {
			return nil
		}
		maxIndex := getMax(nums[left:right])
		leftArrayLeftIndex := left
		leftArrayRightIndex := maxIndex
		rightArrayLeftIndex := maxIndex + 1
		rightArrayrightIndex := right
		newNode := &TreeNode{}
		newNode.Val = nums[maxIndex]
		newNode.Left = traversal(leftArrayLeftIndex, leftArrayRightIndex)
		newNode.Right = traversal(rightArrayLeftIndex, rightArrayrightIndex)
		return newNode
	}
	return traversal(0, len(nums))
}
func Handle6() {
	fmt.Println(constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}))
}
