package list

import "fmt"

// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newNode := &ListNode{
		Next: head,
	}
	// 0->1->2->3->4
	// 0->2->1->3->4
	current := newNode
	for current.Next != nil && current.Next.Next != nil {
		next1 := current.Next
		next2 := current.Next.Next
		next3 := next2.Next

		current.Next = next2
		current.Next.Next = next1
		current.Next.Next.Next = next3

		current = current.Next.Next
	}
	return newNode.Next
}

func Handle5() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	fmt.Println(swapPairs(node))
}
