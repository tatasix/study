package hashtable

import "fmt"

// 349. 两个数组的交集
// 力扣题目链接(opens new window)
//
// 题意：给定两个数组，编写一个函数来计算它们的交集。
func intersection1(nums1 []int, nums2 []int) []int {
	count1 := make([]int, 1001, 1001)
	count2 := make([]int, 1001, 1001)
	res := make([]int, 0)
	for _, v := range nums1 {
		count1[v] = 1
	}
	for _, v := range nums2 {
		count2[v] = 1
	}
	for i := 0; i <= 1000; i++ {
		if count1[i]+count2[i] == 2 {
			res = append(res, i)
		}
	}
	return res
}

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	maxLength := len(nums1)
	length2 := len(nums2)
	if maxLength < length2 {
		maxLength = length2
	}
	count1 := make([]int, maxLength+1, maxLength+1)
	count2 := make([]int, maxLength+1, maxLength+1)
	for _, v := range nums1 {
		count1[v] = 1
	}
	for _, v := range nums2 {
		count2[v] = 1
	}
	for i := 0; i <= maxLength; i++ {
		if count1[i]+count2[i] == 2 {
			res = append(res, i)
		}
	}
	return res
}

func intersection3(nums1 []int, nums2 []int) []int {
	var res []int
	length2 := len(nums2)

	set := make(map[int]struct{}, length2)
	for _, v := range nums2 {
		set[v] = struct{}{}
	}

	for _, v := range nums1 {
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}

func Handle3() {
	num1 := []int{4, 9, 5}
	num2 := []int{9, 4, 9, 8, 4}

	res := intersection3(num1, num2)
	fmt.Println(res)
}
