package backtrack

func solveNQueens(n int) [][]string {
	var res [][]string
	var chessBoard [][]string

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessBoard[i][j] = "."
		}
	}
	var backTrack func(row int)
	backTrack = func(row int) {
		if row == n {

		}
	}
	backTrack(0)
	return res
}

func isValid(n, row, col int, chessBoard [][]string) (res bool) {
	//检查列上所有行
	for i := 0; i < row; i++ {
		if chessBoard[i][col] == "Q" {
			return false
		}
	}
	//检查行上所有列 放的时候，每一行都只会放一个，所以不用检查

	//检查45度角
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessBoard[i][j] == "Q" {
			return false
		}
	}

	//检查135度角
	for i, j := row-1, col+1; j < n && i >= 0; i, j = i-1, j+1 {
		if chessBoard[i][j] == "Q" {
			return false
		}
	}
	return true
}
