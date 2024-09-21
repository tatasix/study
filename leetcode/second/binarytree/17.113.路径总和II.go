package binarytree

func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var path []int
	var count int
	var backTrack func(root *TreeNode)
	backTrack = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		count += root.Val
		if root.Right == nil && root.Left == nil && count == sum {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		} else {
			backTrack(root.Left)
			backTrack(root.Right)
		}
		path = path[:len(path)-1]
		count -= root.Val
	}
	backTrack(root)
	return res
}
