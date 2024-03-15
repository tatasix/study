package greedy

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	l := len(intervals)
	result := make([][]int, 0, l)
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < l; i++ {
		if right >= intervals[i][0] {
			right = max(intervals[i][1], right)
		} else {
			result = append(result, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		}
	}
	result = append(result, []int{left, right})
	return result
}

func Handle20() {
	//[[1,4],[0,2],[3,5]]
	//{ {0, 2},{1, 4}, {3, 5}}
	fmt.Println(merge([][]int{{1, 4}, {0, 2}, {3, 5}}))
}
