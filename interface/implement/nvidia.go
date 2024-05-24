package implement

import "fmt"

type Nvidia struct {
}

func (n *Nvidia) Display(str string) {
	fmt.Println("Nvidia Display " + str)
}
