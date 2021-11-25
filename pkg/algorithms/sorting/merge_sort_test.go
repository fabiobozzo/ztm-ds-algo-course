package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	a := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}

	assert.Equal(t, []int{0, 1, 2, 4, 5, 6, 44, 63, 87, 99, 283}, mergeSort(a))
}
