package slice

import (
	"fmt"
)

func Handle() {
	t13()
}
func t1() {
	s := make([]int, 10)
	s = append(s, 10)
	fmt.Printf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
	// [0 0 0 0 0 0 0 0 0 10] 11 20
	// 双倍扩容
}
func t2() {
	s := make([]int, 0, 10)
	s = append(s, 10)
	fmt.Printf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
	// [10] 1 10 不扩容
}
func t3() {
	s := make([]int, 10, 11)
	s = append(s, 10)
	fmt.Printf("s: %v, len of s: %d, cap of s: %d \n", s, len(s), cap(s))
	// [0, 0, 0, 0, 0, 0, 0, 0, 0, 10] 11 11
}
func t4() {
	s := make([]int, 10, 12)
	s1 := s[8:]
	fmt.Printf("s1: %v, len of s1: %d, cap of s1: %d \n", s1, len(s1), cap(s1))
	// [0,0] 2 4
}
func t5() {
	s := make([]int, 10, 12)
	s1 := s[8:9]
	fmt.Printf("s1: %v, len of s1: %d, cap of s1: %d \n", s1, len(s1), cap(s1))
	// [0] 1 4
}
func t6() {
	s := make([]int, 10, 12)
	s1 := s[8:]
	s1[0] = -1
	fmt.Printf("s: %v", s)
	// [0, 0, 0, 0, 0, 0, 0, 0, -1, 0]
}
func t7() {
	s := make([]int, 10, 12)
	v := s[10]
	// 求问，此时数组访问是否会越界  是
	fmt.Println(v)
}
func t8() {
	s := make([]int, 10, 12)
	s1 := s[8:] // [0,0] l =2 c=4
	s1 = append(s1, []int{10, 11, 12}...)
	fmt.Printf("s1: %v, len of s1: %d, cap of s1: %d \n", s1, len(s1), cap(s1))
	// [0,0,10,11,12] l =5 c=8
	fmt.Printf("s: %v, len of s: %d, cap of s: %d \n", s, len(s), cap(s))
	// [0, 0, 0, 0, 0, 0, 0, 0, 0, 0] l =10 c=12

}
func t9() {
	s := make([]int, 10, 12)
	s1 := s[8:]
	changeSlice(s1)
	fmt.Printf("s: %v", s)
	//[0, 0, 0, 0, 0, 0, 0, 0, -1, 0
}
func changeSlice(s1 []int) {
	s1[0] = -1
}
func t10() {
	s := make([]int, 10, 12)
	s1 := s[8:]
	changeSlice1(s1)
	fmt.Printf("s: %v, len of s: %d, cap of s: %d \n", s, len(s), cap(s))       // 10 12
	fmt.Printf("s1: %v, len of s1: %d, cap of s1: %d \n", s1, len(s1), cap(s1)) // 2 4
}
func changeSlice1(s1 []int) {
	s1 = append(s1, 10)
	fmt.Printf("new s1: %v, len of new s1: %d, cap of new s1: %d \n", s1, len(s1), cap(s1)) // 0 0 10 3 4

}

func t11() {
	s := []int{0, 1, 2, 3, 4}
	a := s[:2]
	b := s[3:]
	fmt.Printf("b: %v, len: %d, cap: %d \n", b, len(b), cap(b)) // [3,4] 2 2
	fmt.Printf("a: %v, len: %d, cap: %d \n", a, len(a), cap(a)) // [0,1] 2 5
	s = append(a, b...)
	fmt.Printf("s: %v, len: %d, cap: %d \n", s, len(s), cap(s)) // [0 1 3 4] 4 5
	fmt.Printf("a: %v, len: %d, cap: %d \n", a, len(a), cap(a)) // [0,1] 2 5
	fmt.Printf("b: %v, len: %d, cap: %d \n", b, len(b), cap(b)) // [4,4] 2 2
	//v := s[4]
	// 是否会数组访问越界
	//fmt.Println(v)
}

func t12() {
	s := make([]int, 512)
	s = append(s, 1)
	fmt.Printf("len of s: %d, cap of s: %d", len(s), cap(s))
}

// 知识点：操作符 [i,j]。基于数组（切片）可以使用操作符 [i,j] 创建新的切片，从索引 i，到索引 j 结束，
// 截取已有数组（切片）的任意部分，返回新的切片，新切片的值包含原数组（切片）的 i 索引的值，
// 但是不包含 j 索引的值。i、j 都是可选的，i 如果省略，默认是 0，j 如果省略，默认是原数组（切片）的长度。i、j 都不能超过这个长度值。
// 假如底层数组的大小为 k，截取之后获得的切片的长度和容量的计算方法：长度：j-i，容量：k-i。
// 截取操作符还可以有第三个参数，形如 [i,j,k]，第三个参数 k 用来限制新切片的容量，
// 但不能超过原数组（切片）的底层数组大小。截取获得的切片的长度和容量分别是：j-i、k-i。
// 所以例子中，切片 t 为 [4]，长度和容量都是 1。
func t13() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t, len(t), cap(t))
}
