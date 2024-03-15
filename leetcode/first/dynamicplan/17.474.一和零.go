package dynamicplan

import "fmt"

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < len(strs); i++ {
		zero, one := 0, 0
		for _, v := range strs[i] {
			if v == '0' {
				zero++
			}
		}
		one = len(strs[i]) - zero

		for j := m; j >= zero; j-- {
			for k := n; k >= one; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zero][k-one]+1)
			}
		}
	}

	return dp[m][n]
}

func Handle17() {
	str := []string{"10", "0", "1"}
	m, n := 1, 1
	fmt.Println(findMaxForm(str, m, n))
	a := `
	物品：字符串数组
	背包：这个背包有两个维度，m(个0),n(个1)
	
	初始化

`
	fmt.Println(a)

}
