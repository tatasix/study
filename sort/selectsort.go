package sort

func SelectSort(arr []int) {
	l := len(arr)
	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[min], arr[i] = arr[i], arr[min]
		}
	}
}
