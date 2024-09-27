package greedy

import "fmt"

//在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
//你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
//给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
// 示例 1: 输入: gas = [1,2,3,4,5],
//             cost = [3,4,5,1,2] 输出: 3
// 示例 2: 输入: gas = [2,3,4], cost = [3,4,3] 输出: -1

func canCompleteCircuit(gas, cost []int) int {
	l := len(gas)
	var remain int
	var total int
	var index int
	for i := 0; i < l; i++ {
		difference := gas[i] - cost[i]
		total += difference
		if remain == 0 && difference > 0 {
			index = i
		}
		remain += difference
		if remain < 0 {
			remain = 0
		}
	}
	if total < 0 {
		return -1
	}
	return index
}

func Handle11() {
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}
