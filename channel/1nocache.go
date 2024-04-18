package channel

import (
	"fmt"
	"time"
)

func noCache() {
	var ch chan int
	ch = make(chan int)
	go func() {
		res := <-ch
		fmt.Println(res)
	}()
	ch <- 10
	time.Sleep(2 * time.Second)
}
func Handle11() {
	noCache()
}
