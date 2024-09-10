package stackqueue

//请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
//实现 MyQueue 类：
//
//void push(int x) 将元素 x 推到队列的末尾
//int pop() 从队列的开头移除并返回元素
//int peek() 返回队列开头的元素
//boolean empty() 如果队列为空，返回 true ；否则，返回 false
//说明：
//你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
//你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
//输入：
//["MyQueue", "push", "push", "peek", "pop", "empty"]
//[[], [1], [2], [], [], []]
//输出：
//[null, null, null, 1, 1, false]

type MyQueue struct {
	In  []int
	Out []int
}

func Constructor2() MyQueue {
	return MyQueue{
		In:  make([]int, 0),
		Out: make([]int, 0),
	}
}

func (m *MyQueue) Push(x int) {
	m.In = append(m.In, x)
}

// Pop 从队列的开头移除并返回元素
func (m *MyQueue) Pop() int {
	var out int
	if len(m.Out) == 0 {
		for i := len(m.In) - 1; i >= 0; i-- {
			temp := m.In[i]
			m.In = m.In[:i]
			m.Out = append(m.Out, temp)
		}
	}
	o := len(m.Out)
	if o > 0 {
		out = m.Out[o-1]
		m.Out = m.Out[:o-1]
	}
	return out
}

// Peek 返回队列开头的元素
func (m *MyQueue) Peek() int {
	var out int
	if len(m.Out) == 0 {
		for i := len(m.In) - 1; i >= 0; i-- {
			temp := m.In[i]
			m.In = m.In[:i]
			m.Out = append(m.Out, temp)
		}
	}
	o := len(m.Out)
	if o > 0 {
		out = m.Out[o-1]
	}
	return out
}

func (m *MyQueue) Empty() bool {
	return len(m.Out) == 0 && len(m.In) == 0
}

type MyQueue1 struct {
	StackIn  []int
	StackOut []int
}

func Constructor1() MyQueue1 {
	return MyQueue1{
		StackIn:  make([]int, 0),
		StackOut: make([]int, 0),
	}
}

func (this *MyQueue1) Push1(x int) {
	this.StackIn = append(this.StackIn, x)
}

func (this *MyQueue1) Pop1() int {
	inLen, outLen := len(this.StackIn), len(this.StackOut)
	if outLen == 0 {
		if inLen == 0 {
			return -1
		}
		for i := inLen - 1; i >= 0; i-- {
			this.StackOut = append(this.StackOut, this.StackIn[i])
		}
		this.StackIn = []int{}
		outLen = len(this.StackOut)
	}
	val := this.StackOut[outLen-1]
	this.StackOut = this.StackOut[:outLen-1]
	return val
}

func (this *MyQueue1) Peek1() int {
	val := this.Pop1()
	if val == -1 {
		return -1
	}
	this.StackOut = append(this.StackOut, val)
	return val
}

func (this *MyQueue1) Empty1() bool {
	return len(this.StackIn) == 0 && len(this.StackOut) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
