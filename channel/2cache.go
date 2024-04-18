package channel

import (
	"fmt"
	"time"
)

func cache() {
	var ch chan int
	ch = make(chan int, 10)
	ch <- 10
	go func() {
		res := <-ch
		fmt.Println(res)
	}()
	//ch <- 10
	time.Sleep(2 * time.Second)
}
func Handle22() {
	cache()
}
