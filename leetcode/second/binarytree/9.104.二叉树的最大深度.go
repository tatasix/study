package binarytree

//[3,9,20,null,null,15,7]

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
