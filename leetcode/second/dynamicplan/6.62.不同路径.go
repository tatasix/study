package dynamicplan

import "fmt"

//一个机器人位于一个 m x n 网格的左上⻆ (起始点在下图中标记为 “Start” )。
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下⻆(在下图中标记为 “Finish” )。
//问总共有多少条不同的路径?

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

//输入:m = 3, n = 7 输出:28
//示例 2: 输入:m = 2, n = 3 输出:3
//示例 3: 输入:m = 7, n = 3 输出:28
//示例 4: 输入:m = 3, n = 3 输出:6

func Handle6() {
	fmt.Println(uniquePaths(7, 3))
}
