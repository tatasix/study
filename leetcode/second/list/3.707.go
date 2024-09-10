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

type ListNode struct {
	Val  int
	Next *ListNode
}
type MyLinkedList struct {
	dummyHead *ListNode // 虚拟头节点
	Size      int       // 链表大小
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		dummyHead: &ListNode{
			Val:  0,
			Next: nil,
		},
		Size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || this == nil || index > this.Size {
		return -1
	}
	current := this.dummyHead.Next
	for i := 0; i < index; i++ {
		if current == nil || current.Next == nil {
			return -1
		}
		current = current.Next
	}
	return current.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &ListNode{
		Val:  val,
		Next: this.dummyHead.Next,
	}
	this.dummyHead = newNode
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &ListNode{
		Val:  val,
		Next: nil,
	}

	current := this.dummyHead

	for current != nil {
		if current.Next == nil {
			current.Next = newNode
		} else {
			current = current.Next
		}
	}
	this.Size++
}

// AddAtIndex 在链表中的第 index 个节点之前添加值为 val  的节点。
// 如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	}

	if index < 0 {
		newNode := &ListNode{
			Val:  val,
			Next: this.dummyHead.Next,
		}
		this.dummyHead = newNode
		this.Size++
		return
	}
	current := this.dummyHead

	for i := 0; i < index; i++ {
		current = current.Next
	}
	newNode := &ListNode{
		Val:  val,
		Next: current.Next.Next,
	}
	current.Next = newNode
	this.Size++
}

// DeleteAtIndex 如果索引 index 有效，则删除链表中的第 index 个节点。
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 && index > this.Size {
		return
	}
	current := this.dummyHead

	for i := 0; i < index; i++ {
		current = current.Next
	}
	current = current.Next.Next
	this.Size--
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
