package dynamicplan

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func rob(root *TreeNode) int {
	// ⻓度为2的数组，0:不偷，1:偷
	var backTrack func(*TreeNode) []int
	backTrack = func(node *TreeNode) []int {
		if node == nil {
			return []int{0, 0}
		}
		left := backTrack(node.Left)
		right := backTrack(node.Right)
		val1 := node.Val + left[0] + right[0] //偷
		val2 := max(left[0], left[1]) + max(right[0], right[1])
		return []int{val2, val1}
	}
	res := backTrack(root)
	return max(res[0], res[1])
}
