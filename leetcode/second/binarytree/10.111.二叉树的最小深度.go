package binarytree

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	} else if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return Min(minDepth(root.Left), minDepth(root.Right)) + 1
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
