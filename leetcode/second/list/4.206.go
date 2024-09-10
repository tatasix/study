package list

import "fmt"

//206.反转链表
//题意：反转一个单链表。
//示例: 输入: 1->2->3->4->5->NULL 输出: 5->4->3->2->1->NULL

func reverseList(node *ListNode) {
	dummyHead := &ListNode{
		Val:  0,
		Next: node,
	}
	current := dummyHead
	var pre *ListNode
	for current != nil && current.Next != nil {
		pre = current.Next
		current.Next = current.Next.Next
		current.Next.Next = pre
		current = current.Next
	}
}

func Handle() {
	node := &ListNode{
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
	reverseList(node)
	fmt.Println(node)
}
