package tree

import "fmt"

/**
 * Definition for a binary tree node.
 *
 */

type TreeNode1 struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RightSideView(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	nodes := []*TreeNode{root}

	if len(nodes) > 0 {

		currentLength := len(nodes)
		for i := 0; i < currentLength; i++ {
			current := nodes[0]
			if current.Left != nil {
				nodes = append(nodes, current.Left)
			}
			if current.Right != nil {
				nodes = append(nodes, current.Right)
			}
			if i == currentLength-1 {
				result = append(result, current.Val)
			}
			nodes = nodes[1:]
		}

	}

	return result
}

func Handle1() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(RightSideView(root))
}
