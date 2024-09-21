package binarytree

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func preorderTraversal1(root *TreeNode) []int {
	var res []int
	var backTravel func(root *TreeNode)
	backTravel = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		backTravel(root.Left)
		backTravel(root.Right)
	}
	backTravel(root)
	return res
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)
	for len(nodes) > 0 {
		l := len(nodes)
		current := nodes[l-1]
		nodes = nodes[:l-1]
		res = append(res, current.Val)
		if current.Right != nil {
			nodes = append(nodes, current.Right)
		}
		if current.Left != nil {
			nodes = append(nodes, current.Left)
		}

	}

	return res
}
