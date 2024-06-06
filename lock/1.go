package lock

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.RWMutex
var count int

func Handle() {
	I()
	return
	go A()
	time.Sleep(2 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Println(count)
}
func A() {
	mu.RLock()
	defer mu.RUnlock()
	B()
}
func B() {
	time.Sleep(5 * time.Second)
	C()
}
func C() {
	mu.RLock()
	defer mu.RUnlock()
}

func D() {
	m := map[string]int{
		"G": 7, "A": 1,
		"C": 3, "E": 5,
		"D": 4, "B": 2,
		"F": 6, "I": 9,
		"H": 8,
	}
	var order []string
	for k, _ := range m {
		order = append(order, k)
	}
	fmt.Println(order)
}
func watShadowDefer(i int) (ret int) {
	ret = i * 2
	if ret > 10 {
		ret = 10
		defer func() {
			ret = ret + 1
		}()
	}
	return
}

func E() {
	result := watShadowDefer(50)
	fmt.Println(result)
}

type foo struct{ Val int }

type bar struct{ Val int }

func F() {
	a := &foo{Val: 5}
	b := &foo{Val: 5}
	c := foo{Val: 5}
	d := bar{Val: 5}
	e := bar{Val: 5}
	f := bar{Val: 5}
	fmt.Print(a == b, c == foo(d), e == f)
}

type User struct {
	Name string
}

func (u *User) SetName(name string) {
	u.Name = name
	fmt.Println(u.Name)
}

type Employee User

func G() {
	employee := new(Employee)
	employee.Name = "Jack"
}

type P *int
type Q *int

func H() {
	var p P = new(int)
	*p += 8
	var x *int = p
	var q Q = x
	*q++
	fmt.Println(*p, *q)
}

const (
	info  = "msg"
	bzero = iota
	bone  = iota
)

func I() {
	fmt.Println(bzero, bone)
}
