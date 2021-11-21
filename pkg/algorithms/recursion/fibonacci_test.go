package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacciIterative(t *testing.T) {
	assert.Equal(t, []int{0}, fibonacciIterative(0))
	assert.Equal(t, []int{0}, fibonacciIterative(1))
	assert.Equal(t, []int{0, 1}, fibonacciIterative(2))
	assert.Equal(t, []int{0, 1, 1}, fibonacciIterative(3))
	assert.Equal(t, []int{0, 1, 1, 2}, fibonacciIterative(4))
	assert.Equal(t, []int{0, 1, 1, 2, 3}, fibonacciIterative(5))
	assert.Equal(t, []int{0, 1, 1, 2, 3, 5}, fibonacciIterative(6))
	assert.Equal(t, []int{0, 1, 1, 2, 3, 5, 8}, fibonacciIterative(7))
	assert.Equal(t, []int{0, 1, 1, 2, 3, 5, 8, 13}, fibonacciIterative(8))
	assert.Equal(t, []int{0, 1, 1, 2, 3, 5, 8, 13, 21}, fibonacciIterative(9))
}

func TestFibonacciRecursive(t *testing.T) {
	var series []int
	for i := 0; i < 9; i++ {
		series = append(series, fibonacciRecursive(i))
	}

	assert.Equal(t, []int{0, 1, 1, 2, 3, 5, 8, 13, 21}, series)
}
