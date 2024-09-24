package backtrack

import "fmt"

//编写一个程序，通过填充空格来解决数独问题。
//
//数独的解法需 遵循如下规则：
//
//数字 1-9 在每一行只能出现一次。
//数字 1-9 在每一列只能出现一次。
//数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
//数独部分空格内已填入了数字，空白格用 '.' 表示。

func solveSudoku(board [][]byte) {
	var backTrack func() bool
	backTrack = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != '.' {
					continue
				}
				for k := '1'; k <= '9'; k++ {
					if !isValid21(board, i, j, byte(k)) {
						continue
					}
					board[i][j] = byte(k)
					if backTrack() {
						return true
					}
					board[i][j] = '.'
				}
				return false
			}
		}
		return true
	}
	backTrack()
}

func isValid21(board [][]byte, row, clo int, n byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][clo] == n || board[row][i] == n {
			return false
		}
	}

	i, j := (row/3)*3, (clo/3)*3
	for k := i; k < 3+i; k++ {
		for l := j; l < 3+j; l++ {
			if board[k][l] == n {
				return false
			}
		}
	}
	return true
}

func Handle() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println("bbbb")
	solveSudoku(board)
	fmt.Println(board)
}
