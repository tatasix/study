package binarytree

import "math"

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	rootIndex := getMax(nums)
	return &TreeNode{
		Left:  constructMaximumBinaryTree(nums[:rootIndex]),
		Right: constructMaximumBinaryTree(nums[rootIndex+1:]),
		Val:   nums[rootIndex],
	}
}

func getMax(nums []int) int {
	min := math.MinInt32
	var res int
	for i, v := range nums {
		if v > min {
			min = v
			res = i
		}
	}
	return res
}
