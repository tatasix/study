package binarytree

//给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
//假设二叉树中至少有一个节点。

func findBottomLeftValue(root *TreeNode) int {
	var res int
	var max int
	if root == nil {
		return 0
	}
	var backTrack func(root *TreeNode, depth int)
	backTrack = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil && depth > max {
			res = root.Val
			max = depth
		}
		if root.Left != nil {
			backTrack(root.Left, depth+1)
		}
		if root.Right != nil {
			backTrack(root.Right, depth+1)
		}
	}
	backTrack(root, 1)
	return res
}

//
//var depth int // 全局变量 最大深度
//var res int   // 记录最终结果
//func findBottomLeftValue(root *TreeNode) int {
//	depth, res = 0, 0 // 初始化
//	dfs(root, 1)
//	return res
//}
//
//func dfs(root *TreeNode, d int) {
//	if root == nil {
//		return
//	}
//	// 因为先遍历左边，所以左边如果有值，右边的同层不会更新结果
//	if root.Left == nil && root.Right == nil && depth < d {
//		depth = d
//		res = root.Val
//	}
//	dfs(root.Left, d+1) // 隐藏回溯
//	dfs(root.Right, d+1)
//}
