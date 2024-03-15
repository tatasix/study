package tree

type Node1 struct {
	Val   int
	Left  *Node1
	Right *Node1
	Next  *Node1
}

func connect(root *Node1) *Node1 {
	if root == nil {
		return nil
	}

	nodes := []*Node1{root}

	for len(nodes) > 0 {
		length := len(nodes)
		//var pre *Node1
		for i := 0; i < length; i++ {
			current := nodes[0]
			nodes = nodes[1:]
			if current.Left != nil {
				nodes = append(nodes, current.Left)
			}
			if current.Right != nil {
				nodes = append(nodes, current.Right)
			}
			if i != length-1 {
				current.Next = nodes[1]
			}
		}
	}

	return root
}

//
//func connect(root *Node) *Node {
//	if root == nil { //防止为空
//		return root
//	}
//	queue := list.New()
//	queue.PushBack(root)
//	tmpArr := make([]*Node, 0)
//	for queue.Len() > 0 {
//		length := queue.Len() //保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
//		for i := 0; i < length; i++ {
//			node := queue.Remove(queue.Front()).(*Node) //出队列
//			if node.Left != nil {
//				queue.PushBack(node.Left)
//			}
//			if node.Right != nil {
//				queue.PushBack(node.Right)
//			}
//			tmpArr = append(tmpArr, node) //将值加入本层切片中
//		}
//		if len(tmpArr) > 1 {
//			// 遍历每层元素,指定next
//			for i := 0; i < len(tmpArr)-1; i++ {
//				tmpArr[i].Next = tmpArr[i+1]
//			}
//		}
//		tmpArr = []*Node{} //清空层的数据
//	}
//	return root
//}
