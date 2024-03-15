package link

import "fmt"

type LinkNode struct {
	Val  int
	Next *LinkNode
}

type MyLinkedList struct {
	Size int
	Head *LinkNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Size: 0,
		Head: &LinkNode{
			Val:  -1000,
			Next: nil,
		},
	}
}

func (this *MyLinkedList) Get(index int) int {
	if this == nil || this.Size < index || index < 0 {
		return -1
	}
	current := this.Head.Next
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &LinkNode{
		Val:  val,
		Next: this.Head.Next,
	}
	this.Head.Next = newNode
	this.Size++
	return
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &LinkNode{
		Val: val,
	}
	current := this.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	this.Size++
	return
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	newNode := &LinkNode{
		Val: val,
	}

	if index > this.Size {
		return
	}
	current := this.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
	this.Size++
	return
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.Size {
		return
	}
	current := this.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	if current.Next != nil && current.Next.Next != nil {
		current.Next = current.Next.Next
	}
	this.Size--
	return
}

// 打印链表
func (list *MyLinkedList) printLinkedList() {
	cur := list.Head      // 设置当前节点为虚拟头节点
	for cur.Next != nil { // 遍历链表
		fmt.Println(cur.Next.Val) // 打印节点值
		cur = cur.Next            // 切换到下一个节点
	}
}

func Handle2() {
	obj := Constructor()
	obj.AddAtHead(100)       // 在头部添加元素
	obj.AddAtTail(242)       // 在尾部添加元素
	obj.AddAtTail(777)       // 在尾部添加元素
	obj.AddAtIndex(1, 99999) // 在指定位置添加元素

	obj.printLinkedList()
	obj.AddAtHead(2)
	obj.AddAtTail(2)
	obj.AddAtIndex(3, 4)
	obj.DeleteAtIndex(5)
}
