package binarytree

// 给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，
// 并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
//
// 一般来说，删除节点可分为两个步骤：
//
// 首先找到需要删除的节点； 如果找到了，删除它。 说明： 要求算法时间复杂度为 $O(h)$，
// h 为树的高度。
func deleteNode(root *TreeNode, key int) *TreeNode {

	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	if root.Left == nil && root.Right == nil {
		return nil
	}
	if root.Left == nil && root.Right != nil {
		return root.Right
	}
	if root.Left != nil && root.Right == nil {
		return root.Left
	}
	current := root.Right
	for current.Left != nil {
		current = current.Left
	}
	current.Left = root.Left
	return root.Right
}

func deleteNode1(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	var traversal func(root *TreeNode) *TreeNode

	traversal = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}

		if root.Val == key {
			if root.Left == nil && root.Right == nil {
				return nil
			} else if root.Left == nil && root.Right != nil {
				return root.Right
			} else if root.Left != nil && root.Right == nil {
				return root.Left
			} else if root.Left != nil && root.Right != nil {
				current := root.Right
				for current.Left != nil {
					current = current.Left
				}
				current.Left = root.Left
				return root.Right
			}
		}
		if root.Val < key {
			root.Right = traversal(root.Right)
		}
		if root.Val > key {
			root.Left = traversal(root.Left)
		}
		return root
	}
	return traversal(root)
}
