package _defer

import "fmt"

func f11() {
	defer func() {
		fmt.Println(recover()) // 3
	}()

	defer func() {
		panic(3) // 将替换恐慌2
	}()
	defer panic(2) // 将替换恐慌1
	defer panic(1) // 将替换恐慌0
	panic(0)
}

func Handle() {
	f21()
}

func f12() {

	// 新开辟一个协程。
	go func() {
		// 一个匿名函数调用。
		// 当它退出完毕时，恐慌2将传播到此新协程的入口
		// 调用中，并且替换掉恐慌0。恐慌2永不会被恢复。
		defer func() {
			// 上一个例子中已经解释过了：恐慌2将替换恐慌1.
			defer panic(2)

			// 当此匿名函数调用退出完毕后，恐慌1将传播到刚
			// 提到的外层匿名函数调用中并与之关联起来。
			func() {
				panic(1)
				// 在恐慌1产生后，此新开辟的协程中将共存
				// 两个未被恢复的恐慌。其中一个（恐慌0）
				// 和此协程的入口函数调用相关联；另一个
				// （恐慌1）和当前这个匿名调用相关联。
			}()
		}()
		panic(0)
	}()

	select {}

}
func f13() {
	defer func() {
		defer func() {
			recover() // 空操作
		}()
	}()
	defer func() {
		func() {
			recover() // 空操作
		}()
	}()
	func() {
		defer func() {
			recover() // 空操作
		}()
	}()
	func() {
		defer recover() // 空操作
	}()
	func() {
		recover() // 空操作
	}()
	recover() // 空操作
	defer func() {
		recover()
	}()
	defer recover() // 空操作
	panic("bye")
}
func f14() {
	defer func() {
		defer func() {
			fmt.Printf("bbbb:%+v \n", recover()) // 此调用将恢复恐慌2
		}()
		fmt.Printf("aaaa:%+v \n", recover())
		defer fmt.Printf("cccc:%+v \n", recover()) // 空操作

		panic(2)
	}()
	panic(1)
}
func f15() {
	defer func() {
		fmt.Printf("cccc:%+v \n", recover())
	}()
	defer func() {
		defer fmt.Printf("aaaa:%+v \n", recover())
		panic(1)
	}()
	defer fmt.Printf("bbbb:%+v \n", recover())
	panic(2)
}
func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}

func f16() {
	f := F(5)
	defer func() {
		fmt.Println(f())
	}()
	defer fmt.Println(f())
	i := f()
	fmt.Println(i)
}
func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()
	f = func() {
		r += 2
	}
	defer f()

	return n + 1
}

func f17() {
	fmt.Println(f(3))
}
func f18() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

func f19() {
	type aa struct {
		Name string
		Age  int
	}

	type bb struct {
		Name string
	}
	a := aa{"aa", 0}
	b := aa{"bb", 1}
	if a == b {
		fmt.Println(1)
	} else {
		fmt.Println(2)
	}
}

func f20() {
	if a := 1; a != 1 {
		fmt.Println("a is 1")
	} else {
		fmt.Println(a)
	}
}

type People struct{}

func (p People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func f21() {
	t := &Teacher{}
	t.ShowB()
}
