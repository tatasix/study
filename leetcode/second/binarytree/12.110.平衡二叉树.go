package binarytree

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		r, l := dfs(root.Right), dfs(root.Left)
		if r == -1 || l == -1 || r-l > 1 || l-r > 1 {
			return -1
		}
		return Max(r, l) + 1
	}
	res := dfs(root)
	return res != -1
}
