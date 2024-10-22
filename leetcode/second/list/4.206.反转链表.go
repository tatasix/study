package list

import (
	"fmt"
)

//206.反转链表
//题意：反转一个单链表。
//示例: 输入: 1->2->3->4->5->NULL 输出: 5->4->3->2->1->NULL

func reverseList(node *ListNode) {

	pre := &ListNode{}
	current := node
	for current != nil {
		temp := current.Next
		current.Next = pre
		pre = current
		current = temp
	}

	return
}
func Handle4() {
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
