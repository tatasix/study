package binarytree

import "math"

//给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。

func getMinimumDifference(root *TreeNode) int {

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
	min := math.MaxInt32
	for i := 1; i < len(result); i++ {
		if result[i]-result[i-1] < min {
			min = result[i] - result[i-1]
		}
	}
	return min
}
