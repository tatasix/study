package binarytree

func buildTree1(inorder []int, postorder []int) *TreeNode {
	li, lp := len(inorder), len(postorder)
	if li == 0 || lp == 0 || li != lp {
		return nil
	}
	top := postorder[lp-1]
	topIndex := find(inorder, top)
	leftInorder := inorder[:topIndex]
	rightInorder := inorder[topIndex+1:]
	return &TreeNode{
		Left:  buildTree(leftInorder, postorder[:len(leftInorder)]),
		Right: buildTree(rightInorder, postorder[len(leftInorder):lp-1]),
		Val:   top,
	}
}

func find(inorder []int, target int) int {
	for i, v := range inorder {
		if v == target {
			return i
		}
	}
	return -1
}
