package dynamicplan

import "fmt"

func findMaxForm1(strs []string, m int, n int) int {
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
	str := []string{"10", "0001", "111001", "1", "0"}
	m, n := 5, 3
	fmt.Println(findMaxForm(str, m, n))

}

// 给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
// 请你找出并返回 strs 的最大子集的大小，该子集中 最多 有 m 个 0 和 n 个 1 。
// 如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。
// 示例 1：
// 输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
// 输出：4
// 解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。 其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 n 的值 3 。
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, v := range strs {
		zero, one := 0, 0
		for _, w := range v {
			if w == '0' {
				zero++
			} else {
				one++
			}
		}
		for j := m; j >= zero; j-- {
			for k := n; k >= one; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zero][k-one]+1)
			}
		}
	}
	return dp[m][n]
}
