package greedy

import (
	"fmt"
	"sort"
)

//有一些球形气球贴在一堵用 XY 平面表示的墙面上。墙面上的气球记录在整数数组 points ，其中points[i] = [xstart, xend] 表示水平直径在 xstart 和 xend之间的气球。你不知道气球的确切 y 坐标。
//一支弓箭可以沿着 x 轴从不同点 完全垂直 地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足  xstart ≤ x ≤ xend，则该气球会被 引爆 。可以射出的弓箭的数量 没有限制 。 弓箭一旦被射出之后，可以无限地前进。
//给你一个数组 points ，返回引爆所有气球所必须射出的 最小 弓箭数 。
// 示例 1： 输入：points = [[10,16],[2,8],[1,6],[7,12]] 输出：2
//示例 2： 输入：points = [[1,2],[3,4],[5,6],[7,8]] 输出：4
//示例 3 输入：points = [[1,2],[2,3],[3,4],[4,5]] 输出：2

func findMinArrowShots(points [][]int) int {
	res := 1

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] {
			res++
		} else {
			if points[i][1] > points[i-1][1] {
				points[i][1] = points[i-1][1]
			}
		}
	}
	return res
}

func Handle17() {
	// {10, 16}, {2, 8}, {1, 6}, {7, 12}
	// {1,2},{3,4},{5,6},{7,8}
	// {1,2},{2,3},{3,4},{4,5}
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))
}
