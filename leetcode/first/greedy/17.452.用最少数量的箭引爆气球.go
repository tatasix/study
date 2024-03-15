package greedy

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	result := 1

	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] {
			result++
		} else {
			points[i][1] = min(points[i][1], points[i-1][1])
		}
	}

	return result

}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
