package sort

import "fmt"

func Handle() {
	origin := []int{1, 2, 0, 22, 4, 3, 12, 9}
	Quicksort(origin)
	//BubbleSort(origin)
	fmt.Println(origin)
}
