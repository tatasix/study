package link

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n < 0 {
		return head
	}
	newHead := &ListNode{
		Next: head,
	}
	slow := newHead
	fast := newHead.Next
	i := 1
	for fast != nil {
		if i > n {
			slow = slow.Next
		}
		fast = fast.Next
		i++
	}
	slow.Next = slow.Next.Next
	return newHead.Next
}
func Handle6() {
}
