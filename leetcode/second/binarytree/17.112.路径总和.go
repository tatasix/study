package binarytree

// 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目
// 标和。
// 说明: 叶子节点是指没有子节点的节点。
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var backTrack func(root *TreeNode, sum int) bool
	backTrack = func(root *TreeNode, sum int) bool {
		if root == nil {
			return false
		}
		current := root.Val + sum
		if root.Left == nil && root.Right == nil {
			return current == targetSum
		}
		return backTrack(root.Left, current) || backTrack(root.Right, current)
	}

	return backTrack(root, 0)
}
