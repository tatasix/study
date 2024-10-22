package list

import "fmt"

// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
// 题目数据 保证 整个链式结构中不存在环。
// 注意，函数返回结果后，链表必须 保持其原始结构 。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	currentA := headA
	for currentA != nil {
		currentB := headB
		for currentB != nil {
			if currentA == currentB {
				return currentA
			}
			currentB = currentB.Next
		}
		currentA = currentA.Next
	}
	return nil
}

func Handle7() {
	node1 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	node2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Println(getIntersectionNode(node1, node2))
}
