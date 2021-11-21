package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSortedArrays(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7}, mergeSortedIntArrays([]int{0, 4, 5, 7}, []int{1, 2, 3, 6}))
	assert.Equal(t, []int{0, 1, 1, 2, 2, 3, 3, 4, 5}, mergeSortedIntArrays([]int{0, 1, 2, 3, 4}, []int{1, 2, 3, 5}))
	assert.Equal(t, []int{0, 1, 1, 2, 2, 3, 3, 4, 5}, mergeSortedIntArrays([]int{1, 2, 3, 5}, []int{0, 1, 2, 3, 4}))
	assert.Equal(t, []int{0, 1, 2, 3}, mergeSortedIntArrays([]int{0, 1, 2, 3}, []int{}))
	assert.Equal(t, []int{0, 1, 1, 1, 2, 3, 3, 5, 6, 8, 10}, mergeSortedIntArrays([]int{0, 1, 2, 3}, []int{1, 1, 3, 5, 6, 8, 10}))
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, mergeSortedIntArrays([]int{1, 3, 6, 8}, []int{0, 2, 4, 5, 7}))
	assert.Equal(t, []int{}, mergeSortedIntArrays([]int{}, []int{}))
}
