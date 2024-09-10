package stackqueue

import "fmt"

// 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。
// 滑动窗口每次只向右移动一位。
// 返回滑动窗口中的最大值。
// 进阶：
// 你能在线性时间复杂度内解决此题吗？
func maxSlidingWindow(nums []int, k int) []int {

	queue := NewMyQueue()
	var res []int
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	res = append(res, queue.Front())
	for i := k; i < len(nums); i++ {
		queue.Push(nums[i])
		queue.Pop(nums[i-k])
		res = append(res, queue.Front())
	}

	return res
}

type MyQueue3 struct {
	e []int
}

func NewMyQueue() *MyQueue3 {
	return &MyQueue3{
		e: make([]int, 0),
	}
}

func (m *MyQueue3) Front() int {
	return m.e[0]
}

func (m *MyQueue3) Back() int {
	return m.e[len(m.e)-1]
}

func (m *MyQueue3) Empty() bool {
	return len(m.e) == 0
}

func (m *MyQueue3) Push(val int) {
	for !m.Empty() && val > m.Back() {
		m.e = m.e[:len(m.e)-1]
	}
	m.e = append(m.e, val)
}

func (m *MyQueue3) Pop(val int) {
	if !m.Empty() && val == m.Front() {
		m.e = m.e[1:]
	}
}
func Handle7() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

//
//// 封装单调队列的方式解题
//type MyQueue struct {
//	queue []int
//}
//
//func NewMyQueue() *MyQueue {
//	return &MyQueue{
//		queue: make([]int, 0),
//	}
//}
//
//func (m *MyQueue) Front() int {
//	return m.queue[0]
//}
//
//func (m *MyQueue) Back() int {
//	return m.queue[len(m.queue)-1]
//}
//
//func (m *MyQueue) Empty() bool {
//	return len(m.queue) == 0
//}
//
//func (m *MyQueue) Push(val int) {
//	for !m.Empty() && val > m.Back() {
//		m.queue = m.queue[:len(m.queue)-1]
//	}
//	m.queue = append(m.queue, val)
//}
//
//func (m *MyQueue) Pop(val int) {
//	if !m.Empty() && val == m.Front() {
//		m.queue = m.queue[1:]
//	}
//}
//
//func maxSlidingWindow(nums []int, k int) []int {
//	queue := NewMyQueue()
//	length := len(nums)
//	res := make([]int, 0)
//	// 先将前k个元素放入队列
//	for i := 0; i < k; i++ {
//		queue.Push(nums[i])
//	}
//	// 记录前k个元素的最大值
//	res = append(res, queue.Front())
//
//	for i := k; i < length; i++ {
//		// 滑动窗口移除最前面的元素
//		queue.Pop(nums[i-k])
//		// 滑动窗口添加最后面的元素
//		queue.Push(nums[i])
//		// 记录最大值
//		res = append(res, queue.Front())
//	}
//	return res
//}
