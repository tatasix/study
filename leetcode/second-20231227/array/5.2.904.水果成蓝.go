package array

import "fmt"

func totalFruit(fruits []int) int {
	left, right, max, length := 0, 0, 0, len(fruits)
	l := make(map[int]int, 2)

	for ; right < length; right++ {
		if _, ok := l[fruits[right]]; ok {
			l[fruits[right]]++
		} else {
			for len(l) >= 2 {
				out := fruits[left]
				l[out]--
				if l[out] == 0 {
					delete(l, out)
				}
				left++
			}
			l[fruits[right]] = 1
		}
		if max < right-left+1 {
			max = right - left + 1
		}
	}
	return max
}

func Handle5() {
	res := totalFruit([]int{1, 0, 0, 0, 1, 0, 4, 0, 4})
	fmt.Println(res)
}
