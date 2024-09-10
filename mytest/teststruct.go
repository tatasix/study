package mytest

import "fmt"

type s1 struct {
	Name string
}

type s2 struct {
	Name string
}

func Handle() {
	s11 := s1{Name: "aa"}
	s21 := s1{Name: "aa"}
	//s12 := s2{Name: "aa"}
	//if s11 == s12 { //不能比较
	//	fmt.Println(" same 1")
	//} else {
	//	fmt.Println("not same 1")
	//}
	if s11 == s21 {
		fmt.Println(" same 2")
	} else {
		fmt.Println("not same 2")
	}

}
