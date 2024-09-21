package binarytree

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}
	var backTrack func(root *TreeNode, path string)
	backTrack = func(root *TreeNode, path string) {
		if root == nil {
			return
		}
		path += "->" + strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			res = append(res, path[2:])
			return
		}
		backTrack(root.Left, path)
		backTrack(root.Right, path)
	}
	backTrack(root, "")
	return res
}
