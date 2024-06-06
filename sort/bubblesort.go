package sort

func BubbleSort(arr []int) {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

//1, 2, 0, 22, 4, 3, 12, 9
//0, 2, 1, 22, 4, 3, 12, 9
