package greedy

import "fmt"

//n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
//你需要按照以下要求，给这些孩子分发糖果：
//每个孩子至少分配到 1 个糖果。
//相邻两个孩子评分更高的孩子会获得更多的糖果。
//请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
//示例 1： 输入：ratings = [1,0,2] 输出：5
//示例 2： 输入：ratings = [1,2,2] 输出：4

func candy(ratings []int) int {
	l := len(ratings)
	candys := make([]int, l)
	candys[0] = 1
	for i := 1; i < l; i++ {
		if ratings[i] > ratings[i-1] {
			candys[i] = candys[i-1] + 1
		} else {
			candys[i] = 1
		}
	}
	var res int
	for i := l - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candys[i] = max(candys[i], candys[i+1]+1)
		}
		res += candys[i]
	}

	return res + candys[l-1]
}

func Handle12() {
	fmt.Println(candy([]int{1, 2, 2}))
}
