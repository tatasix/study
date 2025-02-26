package list

// 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
// 不允许修改 链表。
// 链表中节点的数目范围在范围 [0, 104] 内
// -105 <= Node.val <= 105
// pos 的值为 -1 或者链表中的一个有效索引
func detectCycle(head *ListNode) *ListNode {
	left, right := head, head
	for right != nil && right.Next != nil {
		right = right.Next.Next
		left = left.Next
		if right == left {
			//相遇了
			other := head
			for other != right {
				other = other.Next
				right = right.Next
			}
			return other
		}
	}
	return nil
}

func detectCycle1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {

		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {

			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head

		}
	}
	return nil
}
