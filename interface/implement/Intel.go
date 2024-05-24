package implement

import "fmt"

type Intel struct {
}

func NewIntel() *Intel {
	return &Intel{}
}

func (i *Intel) Display(str string) {
	fmt.Println("Intel Display " + str)
	return
}

func (i *Intel) Storage(str string) {
	fmt.Println("Intel Storage " + str)
}

func (i *Intel) Calculate(str string) {
	fmt.Println("Intel Calculate " + str)
}
