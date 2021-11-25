package sorting

import "math/rand"

func quickSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	left, right := 0, len(a)-1
	pivot := rand.Intn(len(a))

	a[right], a[pivot] = a[pivot], a[right]

	for i := 0; i < len(a); i++ {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}
