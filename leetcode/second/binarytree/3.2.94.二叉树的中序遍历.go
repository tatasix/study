package binarytree

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var backTravel func(root *TreeNode)
	backTravel = func(root *TreeNode) {
		if root == nil {
			return
		}
		backTravel(root.Left)
		res = append(res, root.Val)
		backTravel(root.Right)
	}
	backTravel(root)
	return res
}
