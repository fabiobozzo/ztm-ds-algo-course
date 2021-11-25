package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	a := []int{12, 3, 5, 2, 3, 9, 7}

	assert.Equal(t, []int{2, 3, 3, 5, 7, 9, 12}, insertionSort(a))

	a = []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}

	assert.Equal(t, []int{0, 1, 2, 4, 5, 6, 44, 63, 87, 99, 283}, insertionSort(a))
}
