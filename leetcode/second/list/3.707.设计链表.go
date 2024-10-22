package list

//707.设计链表
//题意：
//在链表类中实现这些功能：
//get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
//addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
//addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
//addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。
//如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
//deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

type MyLinkedList struct {
	Size int
	Head *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (m *MyLinkedList) Get(index int) int {
	if index < 0 || index >= m.Size {
		return -1
	}
	current := m.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Val
}

func (m *MyLinkedList) AddAtHead(val int) {
	newNode := &ListNode{
		Val:  val,
		Next: m.Head,
	}
	m.Head = newNode
	m.Size++
}

func (m *MyLinkedList) AddAtTail(val int) {
	if m.Head == nil {
		m.Head = &ListNode{
			Val:  val,
			Next: nil,
		}
		m.Size++
		return
	}
	current := m.Head
	for i := 1; i < m.Size; i++ {
		current = current.Next
	}

	current.Next = &ListNode{
		Val:  val,
		Next: nil,
	}
	m.Size++
}

func (m *MyLinkedList) AddAtIndex(index, val int) {
	if index == 0 {
		m.AddAtHead(val)
		return
	}
	if index > m.Size {
		return
	}
	if index == m.Size {
		m.AddAtTail(val)
		return
	}
	current := m.Head
	for i := 1; i < index; i++ {
		current = current.Next
	}
	newNode := &ListNode{
		Val:  val,
		Next: current.Next,
	}
	current.Next = newNode
	m.Size++
}

func (m *MyLinkedList) DeleteAtIndex(index int) {
	if m.Size == 0 || m.Size <= index {
		return
	}
	if index == 0 {
		m.Head = m.Head.Next
		m.Size--
		return
	}
	current := m.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	m.Size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// 1 3
//2 3 3
//["MyLinkedList","addAtHead" 1,"addAtTail" 1 3,"addAtIndex" 123,"get" 2,"deleteAtIndex" 1 3,"get" 3,"get","deleteAtIndex","deleteAtIndex","get","deleteAtIndex","get"]
//   []              [1]           [3]            [1,2]           [1]       [1]                [1]   [3]
