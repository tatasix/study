package greedy

//给定一个二叉树，我们在树的节点上安装摄像头。节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。计算监控树的所有节点所需的最小摄像头数量。

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// 0 未覆盖 1 有摄像头 2 有覆盖
// 1 0   1
// 1 1   2
// 1 2   2
// 2 0   1
// 2 2   0
// 0 0   1
func minCameraCover(root *TreeNode) int {
	var res int
	var backTrack func(root *TreeNode) int
	backTrack = func(root *TreeNode) int {
		if root == nil {
			return 2
		}
		left := backTrack(root.Left)
		right := backTrack(root.Right)

		if left == 0 || right == 0 {
			res++
			return 1
		}
		if left == 1 || right == 1 {
			return 2
		}
		return 0
	}
	if backTrack(root) == 0 {
		res++
	}
	return res
}
