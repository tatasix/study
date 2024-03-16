package greedy

import "sort"

// 给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
// 注意: 可以认为区间的终点总是大于它的起点。 区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
// 示例 1:
//
//	 ##
//		###
//		 ##
//		  ##
//
// 输入: [ [1,2], [2,3], [3,4], [1,3] ]
// 输出: 1
// 解释: 移除 [1,3] 后，剩下的区间没有重叠。
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var count int
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			count++
			intervals[i][1] = min(intervals[i][1], intervals[i-1][1])
		}
	}

	return count
}
