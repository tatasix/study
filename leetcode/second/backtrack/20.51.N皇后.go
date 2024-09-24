package backtrack

import (
	"fmt"
	"strings"
)

// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
// 输入：n = 4
// 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// 解释：如上图所示，4 皇后问题存在两个不同的解法。
// 示例 2：
//
// 输入：n = 1
// 输出：[["Q"]]
func solveNQueens(n int) [][]string {
	var res [][]string
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	var backTrack func(board [][]string, row int)
	backTrack = func(board [][]string, row int) {
		if n == row {
			res = append(res, boardTostring(board))
			return
		}
		for clo := 0; clo < n; clo++ {
			if isValid20(board, row, clo) {
				board[row][clo] = "Q"
				backTrack(board, row+1)
				board[row][clo] = "."
			}
		}
	}
	backTrack(board, 0)
	return res
}
func boardTostring(board [][]string) []string {
	var res []string
	for _, v := range board {
		res = append(res, strings.Join(v, ""))
	}
	return res
}
func isValid20(board [][]string, row, clo int) bool {

	for i := 0; i < row; i++ {
		if board[i][clo] == "Q" {
			return false
		}
	}
	for i, j := row-1, clo-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}

	for i, j := row-1, clo+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}

	return true
}

func Handle20() {
	fmt.Println(solveNQueens(4))
}

//func solveNQueens(n int) [][]string {
//    chessBoard := make([][]string,n)
//    var res [][]string
//    for j:=0;j<n;j++{
//        chessBoard[j] = make([]string,n)
//    }
//    for i:=0;i<n;i++{
//        for j:=0;j<n;j++{
//            chessBoard[i][j] = "."
//        }
//    }
//
//    var backtrack func(row int)
//
//    backtrack = func(row int){
//        if n==row{
//            temp := make([]string,n)
//
//            for k,v := range chessBoard{
//                temp[k] = strings.Join(v, "")
//            }
//            res = append(res,temp)
//            return
//        }
//
//        for col:=0;col<n;col++{
//            if isValid(n,row,col,chessBoard){
//                chessBoard[row][col] = "Q"
//                backtrack(row+1)
//                chessBoard[row][col] = "."
//            }
//        }
//    }
//    backtrack(0)
//    return res
//}
//
//func isValid(n, row, col int, chessboard [][]string) bool {
//	for i := 0; i < row; i++ {
//		if chessboard[i][col] == "Q" {
//			return false
//		}
//	}
//	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
//		if chessboard[i][j] == "Q" {
//			return false
//		}
//	}
//	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
//		if chessboard[i][j] == "Q" {
//			return false
//		}
//	}
//	return true
//}
