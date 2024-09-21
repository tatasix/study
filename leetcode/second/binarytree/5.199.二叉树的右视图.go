package binarytree

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		l := len(nodes)
		for i := 0; i < l; i++ {
			current := nodes[i]
			if current.Left != nil {
				nodes = append(nodes, current.Left)
			}
			if current.Right != nil {
				nodes = append(nodes, current.Right)
			}
		}
		res = append(res, nodes[l-1].Val)
		nodes = nodes[l:]
	}
	return res
}
