package binarytree

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var nodes []*TreeNode
	nodes = append(nodes, root)
	for len(nodes) > 0 {
		var path []int
		l := len(nodes)
		for i := 0; i < l; i++ {
			currentNode := nodes[i]
			path = append(path, currentNode.Val)
			if currentNode.Left != nil {
				nodes = append(nodes, currentNode.Left)
			}
			if currentNode.Right != nil {
				nodes = append(nodes, currentNode.Right)
			}
		}
		nodes = nodes[l:]
		res = append(res, path)
	}

	return res
}
