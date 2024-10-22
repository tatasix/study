package list

import "fmt"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newNode := &ListNode{
		Next: head,
	}
	current := newNode
	left := newNode
	for current != nil && n >= 0 {
		current = current.Next
		n--
	}

	for current != nil {
		current = current.Next
		left = left.Next
	}
	if left.Next == nil {
		left.Next = nil
	} else {
		left.Next = left.Next.Next
	}
	return newNode.Next
}

func Handle6() {
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
	removeNthFromEnd(node, 2)
	fmt.Println(node)
}
