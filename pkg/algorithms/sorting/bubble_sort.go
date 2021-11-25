package sorting

func bubbleSort(a []int) []int {
	length := len(a)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}
