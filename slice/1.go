package slice

import (
	"fmt"
)

func Handle() {
	t12()
}
func t1() {
	s := make([]int, 10)
	s = append(s, 10)
	fmt.Printf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))

}
func t2() {
	s := make([]int, 0, 10)
	s = append(s, 10)
	fmt.Printf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}
func t3() {
	s := make([]int, 10, 11)
	s = append(s, 10)
	fmt.Printf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}
func t4() {
	s := make([]int, 10, 12)
	s1 := s[8:]
	fmt.Printf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}
func t5() {
	s := make([]int, 10, 12)
	s1 := s[8:9]
	fmt.Printf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}
func t6() {
	s := make([]int, 10, 12)
	s1 := s[8:]
	s1[0] = -1
	fmt.Printf("s: %v", s)
}
func t7() {
	s := make([]int, 10, 12)
	v := s[10]
	// 求问，此时数组访问是否会越界
	fmt.Println(v)
}
func t8() {
	s := make([]int, 10, 12)
	s1 := s[8:] //l =2 c=4
	s1 = append(s1, []int{10, 11, 12}...)
	fmt.Printf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))

	v := s[10]
	// ...
	// 求问，此时数组访问是否会越界
	fmt.Println(v)

}
func t9() {
	s := make([]int, 10, 12)
	s1 := s[8:]
	changeSlice(s1)
	fmt.Printf("s: %v", s)
}
func changeSlice(s1 []int) {
	s1[0] = -1
}
func t10() {
	s := make([]int, 10, 12)
	s1 := s[8:]
	changeSlice1(s1)
	fmt.Printf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
	fmt.Printf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}
func changeSlice1(s1 []int) {
	s1 = append(s1, 10)
}

func t11() {
	s := []int{0, 1, 2, 3, 4}
	s = append(s[:2], s[3:]...)
	fmt.Printf("s: %v, len: %d, cap: %d", s, len(s), cap(s))
	//v := s[4]
	// 是否会数组访问越界
	//fmt.Println(v)
}

func t12() {
	s := make([]int, 512)
	s = append(s, 1)
	fmt.Printf("len of s: %d, cap of s: %d", len(s), cap(s))
}
