package binarytree

func countNodes1(root *TreeNode) int {
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return dfs(root.Left) + dfs(root.Right) + 1
	}

	return dfs(root)
}

func countNodes(root *TreeNode) int {
	var res int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		res++
	}
	dfs(root)
	return res
}
