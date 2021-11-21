package arrays

import "math"

func mergeSortedIntArrays(a1, a2 []int) []int {
	var mergedArray []int

	if len(a1) == 0 {
		return a2
	}

	if len(a2) == 0 {
		return a1
	}

	i1, i2 := 0, 0

	for i1 < len(a1) || i2 < len(a2) {
		first := math.MaxInt32
		second := math.MinInt32

		if i1 < len(a1) {
			first = a1[i1]
		}

		if i2 < len(a2) {
			second = a2[i2]
		}

		if i2 >= len(a2) || first <= second {
			mergedArray = append(mergedArray, first)
			i1++
		} else if i1 >= len(a1) || first > second {
			mergedArray = append(mergedArray, second)
			i2++
		}
	}

	return mergedArray
}
