package greedy

import "fmt"

func canCompleteCircuit(gas, cost []int) int {
	var result int
	current, total := 0, 0

	for i := 0; i < len(gas); i++ {
		current += gas[i] - cost[i]
		total += gas[i] - cost[i]
		if current < 0 {
			result = i + 1
			current = 0
		}
	}
	if total < 0 {
		return -1
	}
	return result
}

//输入: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
//输出: 3
//gas = [2,3,4] cost = [3,4,3]
//输出: -1

func Handle11() {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	fmt.Println(canCompleteCircuit(gas, cost))
}
