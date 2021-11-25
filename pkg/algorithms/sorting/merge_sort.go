package sorting

func mergeSort(a []int) []int {
	length := len(a)

	if length == 1 {
		return a
	}

	half := length / 2
	left := a[:half]
	right := a[half:]

	merged := merge(mergeSort(left), mergeSort(right))

	return merged
}

func merge(left []int, right []int) []int {
	var merged []int

	leftLength := len(left)
	rightLength := len(right)

	if leftLength == 0 {
		return right
	}

	if rightLength == 0 {
		return left
	}

	leftIndex, rightIndex := 0, 0

	for leftIndex < leftLength || rightIndex < rightLength {
		if leftIndex < leftLength && (rightIndex >= rightLength || left[leftIndex] < right[rightIndex]) {
			merged = append(merged, left[leftIndex])
			leftIndex++
		} else {
			merged = append(merged, right[rightIndex])
			rightIndex++
		}
	}

	return merged
}
