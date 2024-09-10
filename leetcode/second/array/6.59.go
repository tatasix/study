package array

// 59.螺旋矩阵II
// 给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，
// 且元素按顺时针顺序螺旋排列的正方形矩阵。
// 示例:
// 输入: 3 输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]
func generateMatrix(n int) [][]int {
	startx, starty := 0, 0
	var loop int = n / 2
	var center int = n / 2
	count := 1
	offset := 1
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	for loop > 0 {
		i, j := startx, starty

		//行数不变 列数在变
		for j = starty; j < n-offset; j++ {
			res[startx][j] = count
			count++
		}
		//列数不变是j 行数变
		for i = startx; i < n-offset; i++ {
			res[i][j] = count
			count++
		}
		//行数不变 i 列数变 j--
		for ; j > starty; j-- {
			res[i][j] = count
			count++
		}
		//列不变 行变
		for ; i > startx; i-- {
			res[i][j] = count
			count++
		}
		startx++
		starty++
		offset++
		loop--
	}
	if n%2 == 1 {
		res[center][center] = n * n
	}
	return res
}
