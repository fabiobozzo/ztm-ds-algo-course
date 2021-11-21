package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecursiveFactorial(t *testing.T) {
	assert.Equal(t, int64(1), RecursiveFactorial(0))
	assert.Equal(t, int64(1), RecursiveFactorial(1))
	assert.Equal(t, int64(2), RecursiveFactorial(2))
	assert.Equal(t, int64(6), RecursiveFactorial(3))
	assert.Equal(t, int64(24), RecursiveFactorial(4))
	assert.Equal(t, int64(120), RecursiveFactorial(5))
	assert.Equal(t, int64(720), RecursiveFactorial(6))
	assert.Equal(t, int64(5040), RecursiveFactorial(7))
}

func TestIterativeFactorial(t *testing.T) {
	assert.Equal(t, int64(1), IterativeFactorial(0))
	assert.Equal(t, int64(1), IterativeFactorial(1))
	assert.Equal(t, int64(2), IterativeFactorial(2))
	assert.Equal(t, int64(6), IterativeFactorial(3))
	assert.Equal(t, int64(24), IterativeFactorial(4))
	assert.Equal(t, int64(120), IterativeFactorial(5))
	assert.Equal(t, int64(720), IterativeFactorial(6))
	assert.Equal(t, int64(5040), IterativeFactorial(7))
}
