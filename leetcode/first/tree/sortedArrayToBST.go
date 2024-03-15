package tree

type TreeNode8 struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	mid := len(nums) / 2
	newNode := &TreeNode{Val: nums[mid]}
	newNode.Left = sortedArrayToBST(nums[:mid])
	newNode.Right = sortedArrayToBST(nums[mid+1:])
	return newNode
}
