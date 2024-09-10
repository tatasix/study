package stackqueue

import "fmt"

type MyStack struct {
	Queue1 []int
	Queue2 []int
}

func Constructor() MyStack {
	return MyStack{
		Queue1: make([]int, 0),
		Queue2: make([]int, 0),
	}
}

func (m *MyStack) Push(x int) {
	if len(m.Queue1) > 0 {
		m.Queue1 = append(m.Queue1, x)
	} else {
		m.Queue2 = append(m.Queue2, x)
	}
}

func (m *MyStack) Pop() int {
	l1 := len(m.Queue1)
	var out int
	if l1 > 0 {
		for i := 0; i < l1-1; i++ {
			m.Queue2 = append(m.Queue2, m.Queue1[i])
		}
		out = m.Queue1[l1-1]
		m.Queue1 = make([]int, 0)
		return out
	}
	l2 := len(m.Queue2)
	for i := 0; i < l2-1; i++ {
		m.Queue1 = append(m.Queue1, m.Queue2[i])
	}
	out = m.Queue2[l2-1]
	m.Queue2 = make([]int, 0)
	return out
}

func (m *MyStack) Top() int {
	out := m.Pop()
	if len(m.Queue1) > 0 {
		m.Queue1 = append(m.Queue1, out)
	} else {
		m.Queue2 = append(m.Queue2, out)
	}
	return out
}

func (m *MyStack) Empty() bool {
	return len(m.Queue1) == 0 && len(m.Queue2) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func Handle3() {
	obj := Constructor()
	obj.Push(1)
	res := obj.Top()
	fmt.Println(res)
	res1 := obj.Empty()
	fmt.Println(res1)
}
