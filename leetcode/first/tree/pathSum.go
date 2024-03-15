package tree

import "fmt"

type TreeNode3 struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Handle3() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:  4,
			Left: nil,
			Right: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	fmt.Println(PathSum(root, 22))
}

func PathSum(root *TreeNode, targetSum int) [][]int {
	var results [][]int
	if root == nil {
		return results
	}
	var traversal func(root *TreeNode, path []int, total int)
	traversal = func(root *TreeNode, path []int, total int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		total -= root.Val

		if root.Left == nil && root.Right == nil && total == 0 {
			newPath := make([]int, len(path))
			copy(newPath, path)
			results = append(results, newPath)
		}
		if root.Left != nil {

			traversal(root.Left, path, total)

		}
		if root.Right != nil {

			traversal(root.Right, path, total)

		}
	}
	traversal(root, []int{}, targetSum)
	return results
}

func PathSum1(root *TreeNode, targetSum int) [][]int {
	var results [][]int
	if root == nil {
		return results
	}
	nodes := []*TreeNode{root}
	paths := [][]int{{root.Val}}
	vals := []int{root.Val}

	for len(nodes) > 0 {
		length := len(nodes)
		for i := 0; i < length; i++ {
			node := nodes[0]
			path := paths[0]
			val := vals[0]
			nodes = nodes[1:]
			paths = paths[1:]
			vals = vals[1:]

			if node.Left == nil && node.Right == nil && val == targetSum {
				results = append(results, path)
			}

			if node.Left != nil {
				nodes = append(nodes, node.Left)
				pathLeft := make([]int, len(path))
				copy(pathLeft, path)
				pathLeft = append(pathLeft, node.Left.Val)
				paths = append(paths, pathLeft)
				vals = append(vals, node.Left.Val+val)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
				pathRight := make([]int, len(path))
				copy(pathRight, path)
				pathRight = append(pathRight, node.Right.Val)
				paths = append(paths, pathRight)
				vals = append(vals, node.Right.Val+val)
			}
		}
	}

	return results
}
