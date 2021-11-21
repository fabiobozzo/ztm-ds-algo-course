package linked_lists

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppendPrepend(t *testing.T) {
	list := NewSinglyLinkedList("b")

	list.Append("c")
	list.Append("d")
	list.Prepend("a")

	assert.Equal(t, "a, b, c, d", fmt.Sprint(list))

	list.Prepend("0")
	list.Append("99")
	assert.Equal(t, "0, a, b, c, d, 99", fmt.Sprint(list))
}

func TestInsertLookup(t *testing.T) {
	list := NewSinglyLinkedList("a")

	list.Append("b")
	list.Append("c")
	list.Append("d")
	list.Append("e")
	list.Append("f")

	element, err := list.Lookup(0)
	assert.NoError(t, err)
	assert.Equal(t, "a", element.Value)

	element, err = list.Lookup(5)
	assert.NoError(t, err)
	assert.Equal(t, "f", element.Value)

	element, err = list.Lookup(2)
	assert.NoError(t, err)
	assert.Equal(t, "c", element.Value)

	err = list.Insert(2, "1")
	assert.NoError(t, err)
	assert.Equal(t, "a, b, 1, c, d, e, f", fmt.Sprint(list))

	err = list.Insert(0, "x")
	assert.NoError(t, err)
	assert.Equal(t, "x, a, b, 1, c, d, e, f", fmt.Sprint(list))

	err = list.Insert(8, "y")
	assert.NoError(t, err)
	assert.Equal(t, "x, a, b, 1, c, d, e, f, y", fmt.Sprint(list))

	err = list.Insert(8, "z")
	assert.NoError(t, err)
	assert.Equal(t, "x, a, b, 1, c, d, e, f, z, y", fmt.Sprint(list))

	err = list.Insert(1, "0")
	assert.NoError(t, err)
	assert.Equal(t, "x, 0, a, b, 1, c, d, e, f, z, y", fmt.Sprint(list))
}

func TestRemove(t *testing.T) {
	list := NewSinglyLinkedList("a")

	list.Append("b")
	list.Append("c")
	list.Append("d")
	list.Append("e")
	list.Append("f")

	err := list.Remove(3)
	assert.NoError(t, err)
	assert.Equal(t, "a, b, c, e, f", fmt.Sprint(list))

	err = list.Remove(0)
	assert.Error(t, err)

	err = list.Remove(5)
	assert.Error(t, err)

	err = list.Remove(4)
	assert.Equal(t, "a, b, c, e", fmt.Sprint(list))
}

func TestReverse(t *testing.T) {
	list := NewSinglyLinkedList("a")

	list.Append("b")
	list.Append("c")
	list.Append("d")
	list.Append("e")
	list.Append("f")

	list.Reverse()

	assert.Equal(t, "f, e, d, c, b, a", fmt.Sprint(list))

	list = NewSinglyLinkedList("a")
	list.Reverse()
	assert.Equal(t, "a", fmt.Sprint(list))
}
