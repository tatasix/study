package greedy

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	var res int
	sort.Ints(g)
	sort.Ints(s)
	j := len(s) - 1
	for i := len(g) - 1; i >= 0; i-- {
		if j >= 0 && g[i] <= s[j] {
			res++
			j--
		}

	}
	return res
}

func Handle2() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{}))
}
