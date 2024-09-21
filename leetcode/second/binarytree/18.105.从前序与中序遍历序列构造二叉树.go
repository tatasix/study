package binarytree

import "fmt"

func buildTree(preorder []int, inorder []int) *TreeNode {
	lp, li := len(preorder), len(inorder)
	if lp == 0 || li == 0 || lp != li {
		return nil
	}

	top := preorder[0]
	topIndex := find(inorder, top)
	leftInorder := inorder[:topIndex]
	rightInorder := inorder[topIndex+1:]

	return &TreeNode{
		Left:  buildTree(preorder[1:len(leftInorder)+1], leftInorder),
		Right: buildTree(preorder[len(leftInorder)+1:], rightInorder),
		Val:   top,
	}
}

func Handle18() {
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
