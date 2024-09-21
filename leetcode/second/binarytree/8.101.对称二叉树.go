package binarytree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var backTrace func(left, right *TreeNode) bool
	backTrace = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		} else if left != nil && right == nil {
			return false
		} else if left == nil && right != nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return backTrace(left.Left, right.Right) && backTrace(left.Right, right.Left)
	}
	return backTrace(root.Left, root.Right)
}
