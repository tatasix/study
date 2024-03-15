package dynamicplan

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
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
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func Handle7() {
	obstacleGrid := [][]int{{0, 0}, {1, 1}, {0, 0}}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}

//[{0,0},{1,1},{0,0}]  0

//obstacleGrid = [{0,0,0},{0,1,0},{0,0,0}]
//输出：2
//obstacleGrid = [{0,1},{0,0}]
//输出：1
