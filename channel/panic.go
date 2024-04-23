package channel

import (
	"fmt"
)

func panic1() {
	c := make(chan int, 0)
	fmt.Println(<-c) //一直接受不到数据，阻塞在这里，死锁
}

func panic2() {
	c := make(chan int, 0)
	c <- 1
	fmt.Println(<-c)
}

// var只定义了类型，但并未初始化，deadlock
func panic3() {
	var c chan int
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}

// 结果是死锁，应该将数据传进channel后，并没有关闭channel，
// for循环接收channel一直在监听，死锁
func panic4() {
	var c chan int        //定义类型
	c = make(chan int, 0) //初始化
	go func() {
		for i := 0; i < 3; i++ {
			c <- i //传入channel数据
		}
		//close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}
func panic5() {
	var c chan int
	close(c)
}

func Handle() {
	panic1()
}

func test() {

}
