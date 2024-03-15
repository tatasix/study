package array

import "fmt"

// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]
func generateMatrix(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	var (
		res    []int
		startX = 0
		startY = 0
		offset = 1
		m      = len(matrix)
		n      = len(matrix[0])
		loop   = (m + 1) / 2
	)
	for loop > 0 {
		x, y := startX, startY

		for ; y < n-offset; y++ {
			res = append(res, matrix[x][y])
		}
		for ; x < m-offset; x++ {
			res = append(res, matrix[x][y])
		}
		for ; y >= offset; y-- {
			res = append(res, matrix[x][y])
		}
		for ; x >= offset; x-- {
			res = append(res, matrix[x][y])
		}
		startX++
		startY++
		offset++
		loop--
	}
	if m == n && m%2 == 1 {
		middle := m / 2
		res = append(res, matrix[middle][middle])
	}
	return res
}

func Handle() {
	res := generateMatrix([][]int{{1, 2, 3, 10}, {4, 5, 6, 11}, {7, 8, 9, 12}})
	fmt.Println(res)
}

//{6, 9, 7}
//{1, 2, 3, 10}, {4, 5, 6, 11}, {7, 8, 9, 12}
// {1,2,3},{4,5,6},{7,8,9}
