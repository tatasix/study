package p2

import (
	"fmt"
)

// PP1 通过接口来避免import PP1 类
// 使用的时候：
// pp1 := p1.PP1{}
// pp1.HelloFromP2Side()
type PP1 interface {
	HelloFromP1()
}

type PP2 struct {
	PP1
}

func New(pp1 PP1) *PP2 {
	return &PP2{
		pp1,
	}
}

func (p *PP2) HelloFromP2() {
	fmt.Println("Hello from package p2")
}

func (p *PP2) HelloFromP1Side() {
	p.PP1.HelloFromP1()
}
