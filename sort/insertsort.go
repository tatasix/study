package sort

func InsertSort(arr []int) {
	l := len(arr)
	for i := 1; i < l; i++ {
		current := arr[i]
		j := i - 1
		for j >= 0 && current < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = current
	}
}
