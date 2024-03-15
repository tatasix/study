package backtrack

func solveSudoku(board [][]byte) {
	var backtrack func() bool

	backtrack = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != '.' {
					continue
				}

				for k := '1'; k <= '9'; k++ {
					if isValid(i, j, byte(k), board) {
						board[i][j] = byte(k)
						if backtrack() {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false

			}
		}
		return true
	}
	backtrack()
}

func isValid(row, col int, v byte, board [][]byte) bool {
	l := len(board)
	for i := 0; i < l; i++ {
		if board[row][i] == v {
			return false
		}
	}

	for i := 0; i < l; i++ {
		if board[i][col] == v {
			return false
		}
	}

	startRow := (row / 3) * 3
	startCol := (col / 3) * 3

	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == v {
				return false
			}
		}
	}
	return true
}
