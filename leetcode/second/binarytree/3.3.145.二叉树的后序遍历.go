package binarytree

func postorderTraversal(root *TreeNode) []int {
	var res []int
	var backTravel func(root *TreeNode)
	backTravel = func(root *TreeNode) {
		if root == nil {
			return
		}
		backTravel(root.Left)
		backTravel(root.Right)
		res = append(res, root.Val)
	}
	backTravel(root)
	return res
}
