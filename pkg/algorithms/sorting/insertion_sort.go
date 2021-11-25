package sorting

func insertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		if a[0] > a[i] {
			temp := a[i]
			copy(a[1:i+1], a[0:i])
			a[0] = temp
		} else {
			for j := 1; j < i; j++ {
				if a[j-1] <= a[i] && a[j] > a[i] {
					a[j], a[i] = a[i], a[j]
				}
			}
		}
	}

	return a
}
