package sorting

func selectionSort(a []int) []int {
	length := len(a)

	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		if min != i {
			a[min], a[i] = a[i], a[min]
		}
	}

	return a
}
