package list

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	pre := &ListNode{}

	for head != nil {
		if head.Val == val {
			pre.Next = head.Next
		}
	}

	return nil
}

//head = [1,2,6,3,4,5,6], val = 6
//输出：[1,2,3,4,5]

func Handel() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	removeElements(head, 6)
}
