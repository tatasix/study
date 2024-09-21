package binarytree

func sumOfLeftLeaves(root *TreeNode) int {
	var res int
	if root == nil {
		return res
	}
	var backTrack func(left *TreeNode, isLeft bool)
	backTrack = func(left *TreeNode, isLeft bool) {
		if left == nil {
			return
		}

		if left.Left == nil && left.Right == nil {
			if isLeft {
				res += left.Val
			}
			return
		} else if left.Left != nil && left.Right == nil {
			backTrack(left.Left, true)
		} else if left.Left == nil && left.Right != nil {
			backTrack(left.Right, false)
		} else {
			backTrack(left.Left, true)
			backTrack(left.Right, false)
		}

	}

	backTrack(root, false)
	return res
}
func sumOfLeftLeaves1(root *TreeNode) int {
	var result int
	if root == nil {
		return result
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		length := len(nodes)
		for i := 0; i < length; i++ {
			current := nodes[0]
			nodes = nodes[1:]
			if current.Left != nil {
				if current.Left.Left == nil && current.Left.Right == nil {
					result = result + current.Left.Val
				}

				nodes = append(nodes, current.Left)
			}
			if current.Right != nil {
				nodes = append(nodes, current.Right)
			}

		}
	}
	return result
}
func sumOfLeftLeaves2(root *TreeNode) int {
	var result int
	if root == nil {
		return result
	}
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {

		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			result += node.Left.Val
		}
		if node.Left != nil {
			travel(node.Left)
		}
		if node.Right != nil {
			travel(node.Right)
		}
	}
	travel(root)
	return result
}
