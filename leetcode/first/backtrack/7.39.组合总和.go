package backtrack

func combinationSum(candidates []int, target int) [][]int {
	var (
		result [][]int
		path   []int
		sum    int
	)
	var backtrack func(index int)
	backtrack = func(index int) {
		if sum > target {
			return
		}
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := index; i < len(candidates); i++ {

			path = append(path, candidates[i])
			sum += candidates[i]
			backtrack(i)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	backtrack(0)
	return result
}
