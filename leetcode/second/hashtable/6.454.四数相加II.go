package hashtable

import "fmt"

//给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，
//使得 A[i] + B[j] + C[k] + D[l] = 0。
//
//为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。
//所有整数的范围在 -2^28 到 2^28 - 1 之间，最终结果不会超过 2^31 - 1 。

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var res int
	set := make(map[int]int)
	for _, v := range nums1 {
		for _, vv := range nums2 {
			set[0-v-vv] += 1
		}
	}
	for _, v := range nums3 {
		for _, vv := range nums4 {
			if vvv, ok := set[v+vv]; ok {
				res += vvv
			}
		}
	}
	return res
}

func Handle6() {
	//res := fourSumCount([]int{0}, []int{0}, []int{0}, []int{0})
	//res := fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2})
	res := fourSumCount([]int{-1, 1}, []int{-1, 1}, []int{-1, 1}, []int{1, -1})
	fmt.Println(res)
}
