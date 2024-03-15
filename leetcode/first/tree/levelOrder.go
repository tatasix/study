package tree

/**
 * Definition for a Node.
 **/

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var results [][]int
	if root == nil {
		return results
	}
	nodes := []*Node{root}
	for len(nodes) > 0 {
		var result []int
		currentLength := len(nodes)
		for i := 0; i < currentLength; i++ {
			current := nodes[0]
			childrenLength := len(current.Children)
			for j := 0; j < childrenLength; j++ {
				if current.Children[j] != nil {
					nodes = append(nodes, current.Children[j])
				}
			}
			result = append(result, current.Val)
			nodes = nodes[1:]
		}
		results = append(results, result)
	}
	return results
}
