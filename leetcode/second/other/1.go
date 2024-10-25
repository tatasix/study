package other

import (
	"fmt"
	"strconv"
)

// 217
// 给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
// 示例 1： 输入：nums = [1,2,3,1] 输出：true 解释： 元素 1 在下标 0 和 3 出现。
// 示例 2： 输入：nums = [1,2,3,4] 输出：false 解释：所有元素都不同。
// 示例 3:  输入：nums = [1,1,1,3,3,4,3,2,4,2] 输出：true
func containsDuplicate(nums []int) bool {
	l := len(nums)
	data := make(map[int]struct{})

	for i := 0; i < l; i++ {
		if _, ok := data[nums[i]]; ok {
			return true
		}
		data[nums[i]] = struct{}{}
	}
	return false
}

// 61 旋转链表
// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
// 示例 1： 输入：head = [1,2,3,4,5], k = 2
// 输出：[4,5,1,2,3]
// 示例 2： 输入：head = [0,1,2], k = 4  输出：[2,0,1]
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// rotateRight 旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// 1. 计算链表长度
	length := 1
	current := head
	for current.Next != nil {
		current = current.Next
		length++
	}

	// 形成环
	current.Next = head

	// 计算实际需要移动的步数
	k = k % length
	if k == 0 {
		// 如果不需要移动，则断开环
		current.Next = nil
		return head
	}

	// 找到新的尾节点
	stepsToNewTail := length - k
	newTail := head
	for i := 0; i < stepsToNewTail-1; i++ {
		newTail = newTail.Next
	}

	// 断开环，形成新的头节点
	newHead := newTail.Next
	newTail.Next = nil

	return newHead
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 && n > 0 {
		if nums1[m-1] >= nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	for n > 0 {
		nums1[n-1] = nums2[n-1]
		n--
	}
}
func Handle1() {
	var res []byte
	ss := "Let's take LeetCode contest"
	s := []byte(ss)
	i, j := 0, 0
	for ; i < len(s); i++ {
		if s[i] == ' ' {
			if i != j {
				a := reverse(s[j:i])
				res = append(res, a...)
				res = append(res, ' ')
				j = i + 1
			}
		}
	}
	if i != j {
		a := reverse(s[j:i])
		res = append(res, a...)
	}
	fmt.Println(string(res))
}

func reverse(s []byte) []byte {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return s
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
