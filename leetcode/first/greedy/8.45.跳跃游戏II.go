package greedy

func jump(nums []int) int {
	var result int
	l := len(nums)
	if l < 2 {
		return result
	}
	current, next := 0, 0 //当前最远覆盖下标，其实不管怎么走都是可以走到的

	for i := 0; i < l; i++ {
		temp := i + nums[i]
		if temp > next {
			next = temp
		}
		if current == i {
			result++
			current = next
		}
		if next >= l-1 {
			break
		}
	}
	return result
}
