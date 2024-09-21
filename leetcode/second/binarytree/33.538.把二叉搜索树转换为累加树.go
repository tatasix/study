package binarytree

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	sum := 0
	var backTrack func(root *TreeNode) *TreeNode
	backTrack = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		root.Right = backTrack(root.Right)
		sum += root.Val
		root.Val = sum
		root.Left = backTrack(root.Left)
		return root
	}
	return backTrack(root)
}
