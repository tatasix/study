package _interface

import (
	"fmt"
	"study/interface/implement"
)

func Handle() {
	res, err := NewMyPlan(implement.NewBiliBili()).Learn("ss")
	fmt.Println(res)
	fmt.Println(err)
}
