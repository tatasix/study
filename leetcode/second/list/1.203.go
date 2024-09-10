package list

//203.移除链表元素
//题意：删除链表中等于给定值 val 的所有节点。
//示例 1： 输入：head = [1,2,6,3,4,5,6], val = 6 输出：[1,2,3,4,5]
//示例 2： 输入：head = [], val = 1 输出：[]
//示例 3： 输入：head = [7,7,7,7], val = 7 输出：[]

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func removeElements1(head *ListNode1, val int) *ListNode1 {

	for head != nil && head.Val == val { //由于leetcode代码运行方式，for循环条件判断前后顺序不能修改，下面的for循环也同样如此
		head = head.Next
	}

	current := head
	for current != nil && current.Next != nil {
		if current.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

func removeElements(head *ListNode1, val int) *ListNode1 {
	dummyHead := &ListNode1{
		Val:  0,
		Next: head,
	}
	current := dummyHead
	for current != nil && current.Next != nil {
		if current.Val == val {
			current = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return dummyHead.Next
}
