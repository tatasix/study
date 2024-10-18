package dynamicplan

import "fmt"

//给定一个mxn的整数数组grid。一个机器人初始位于左上角（即grid[0][0]）
//机器人尝试移动到右下角（即grid[m-1][n-1]）。机器人每次只能向下或者向右移动一步。
//网格中的障碍物和空位置分别用1和0来表示。机器人的移动路径中不能包含任何有障碍物的方格。
//返回机器人能够到达右下角的不同路径数量。
//测试用例保证答案小于等于 2 * 109。

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

//输入：obstacleGrid = {0, 0, 0}, {0, 1, 0}, {0, 0, 0}  输出：2
//输入：obstacleGrid = [[0,1],[0,0]]  输出：1

func Handle7() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))
}
