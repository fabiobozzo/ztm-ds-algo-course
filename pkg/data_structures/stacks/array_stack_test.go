package stacks

import (
	"ds_algo_course/pkg/constants"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayPush(t *testing.T) {
	s := NewArrayStack(10)

	assert.True(t, s.IsEmpty())

	s.Push("a")

	assert.Equal(t, "a", fmt.Sprint(s))
	assert.False(t, s.IsEmpty())

	s.Push("b")
	s.Push("c")

	assert.Equal(t, "c, b, a", fmt.Sprint(s))
}

func TestArrayPopAndPeek(t *testing.T) {
	s := NewArrayStack(10)

	s.Push("a")
	s.Push("b")
	s.Push("c")

	peeked, err := s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "c", peeked)

	popped, err := s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "c", popped)

	peeked, err = s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "b", peeked)

	popped, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "b", popped)

	peeked, err = s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "a", peeked)

	popped, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "a", popped)

	_, err = s.Peek()
	assert.ErrorIs(t, err, constants.ErrEmpty)

	_, err = s.Pop()
	assert.ErrorIs(t, err, constants.ErrEmpty)
}
