package mytest

func ReverseArray(data []int) []int {
	length := len(data)
	for i := 0; i < length/2; i++ {
		data[i], data[length-1-i] = data[length-1-i], data[i]
	}
	return data
}
