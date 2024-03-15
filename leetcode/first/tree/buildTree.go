package tree

import "fmt"

type TreeNode7 struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}
	rootVal := postorder[len(postorder)-1]
	index := find(inorder, rootVal)
	fmt.Println(postorder)
	fmt.Println(inorder)
	fmt.Println(index)
	root := &TreeNode{Val: rootVal}
	root.Left = buildTree(inorder[:index], postorder[:index])

	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])

	return root
}

func find(order []int, value int) (index int) {
	for i := 0; i < len(order); i++ {
		if order[i] == value {
			index = i
			return
		}
	}
	return
}

func Handle7() {
	fmt.Println(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}
