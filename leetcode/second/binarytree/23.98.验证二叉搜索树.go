package binarytree

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	var result []int
	var bfs func(root *TreeNode)
	bfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		bfs(root.Left)
		result = append(result, root.Val)
		bfs(root.Right)
	}
	bfs(root)
	l := len(result)
	if l == 1 {
		return true
	}
	for i := 1; i < l; i++ {
		if result[i-1] >= result[i] {
			return false
		}
	}
	return true
}
