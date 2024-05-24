package _interface

import (
	"fmt"
	"study/interface/implement"
)

func Handle() {
	res, err := NewMyPlan(implement.NewBiliBili()).Learn("ss")
	fmt.Println(res)
	fmt.Println(err)

	NewComputer(implement.NewIntel(), implement.NewIntel(), implement.NewIntel()).DoWork()
	NewComputer(implement.NewIntel(), implement.NewKingston(), &implement.Nvidia{}).DoWork()
}
