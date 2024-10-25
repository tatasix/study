package other

import "fmt"

// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
// 1 2 3
// 4 5 6
// 7 8 9
func spiralOrder(matrix [][]int) []int {
	var result []int
	y := len(matrix)
	if y == 0 {
		return result
	}
	x := len(matrix[0])
	left, right, top, bottom := 0, x-1, 0, y-1

	for left <= right && top <= bottom {
		//从左到右
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		//从上到下
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		//从右到左
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		//从下到上
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}
	return result
}

func Handle54() {
	//{1, 2, 3}, {4, 5, 6}, {7, 8, 9}
	fmt.Println(spiralOrder([][]int{{6, 9, 7}}))
}
