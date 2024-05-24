package implement

import "fmt"

type Kingston struct {
}

func NewKingston() *Kingston {
	return &Kingston{}
}

func (k *Kingston) Storage(str string) {
	fmt.Println("Kingston Storage " + str)
}
