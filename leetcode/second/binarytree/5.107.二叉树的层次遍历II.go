package binarytree

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		l := len(nodes)
		var path []int
		for i := 0; i < l; i++ {
			current := nodes[i]
			if current.Left != nil {
				nodes = append(nodes, current.Left)
			}
			if current.Right != nil {
				nodes = append(nodes, current.Right)
			}
			path = append(path, current.Val)
		}
		nodes = nodes[l:]
		res = append(res, path)
	}
	l := len(res)
	for i := 0; i < l/2; i++ {
		res[i], res[l-1-i] = res[l-1-i], res[i]
	}
	return res
}
