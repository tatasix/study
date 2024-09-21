package binarytree

import "math"

//给定一个有相同值的二叉搜索树(BST)，找出 BST 中的所有众数(出现频率最高的元素)。
//假定 BST 有如下定义:
//  结点左子树中所含结点的值小于等于当前结点的值
//  结点右子树中所含结点的值大于等于当前结点的值
//  左子树和右子树都是二叉搜索树

func findMode1(root *TreeNode) []int {
	res := make(map[int]int)
	var backTrack func(root *TreeNode)
	backTrack = func(root *TreeNode) {
		if root == nil {
			return
		}
		backTrack(root.Left)
		res[root.Val]++
		backTrack(root.Right)
	}
	backTrack(root)
	var result []int
	max := math.MinInt32
	for i, v := range res {
		if v > max {
			max = v
			result = result[:0]
			result = append(result, i)
		} else if v == max {
			result = append(result, i)
		}
	}
	return result
}

func findMode(root *TreeNode) []int {
	var backTrack func(root *TreeNode)
	var result []int
	max := math.MinInt32
	var pre *TreeNode
	var count int
	backTrack = func(root *TreeNode) {
		if root == nil {
			return
		}
		backTrack(root.Left)
		if pre == nil || pre.Val < root.Val {
			count = 1
		} else {
			count++
		}
		if count > max {
			max = count
			result = result[:0]
			result = append(result, root.Val)
		} else if count == max {
			result = append(result, root.Val)
		}
		pre = root
		backTrack(root.Right)
	}
	backTrack(root)

	return result
}
