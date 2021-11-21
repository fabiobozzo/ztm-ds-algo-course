package stacks

import (
	"ds_algo_course/pkg/constants"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPush(t *testing.T) {
	s := NewLinkedListStack()

	assert.True(t, s.IsEmpty())

	s.Push("a")

	assert.Equal(t, "a", fmt.Sprint(s))
	assert.False(t, s.IsEmpty())

	s.Push("b")
	s.Push("c")

	assert.Equal(t, "c, b, a", fmt.Sprint(s))

	concreteType, ok := s.(*LinkedListStack)
	assert.True(t, ok)
	assert.Equal(t, "a", concreteType.Bottom.Value)
	assert.Equal(t, "c", concreteType.Top.Value)
	assert.Equal(t, 3, concreteType.length)
}

func TestPopAndPeek(t *testing.T) {
	s := NewLinkedListStack()

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
