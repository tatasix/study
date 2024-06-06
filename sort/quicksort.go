package sort

func Quicksort(arr []int) {
	quicksort(arr, 0, len(arr)-1)
	return
}

func quicksort(a []int, begin, end int) {
	if begin < end {
		center := end
		counter := begin
		for i := begin; i < end; i++ {
			if a[i] < a[center] {
				//交换
				a[i], a[counter] = a[counter], a[i]
				counter++
			}
		}
		//交换pivot;
		a[center], a[counter] = a[counter], a[center]
		//向下调用;
		quicksort(a, begin, counter-1)
		quicksort(a, counter+1, end)
	}
}
