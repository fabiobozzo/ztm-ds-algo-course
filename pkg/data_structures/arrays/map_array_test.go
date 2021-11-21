package arrays

import (
	"ds_algo_course/pkg/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapArray(t *testing.T) {
	var item Item
	var err error

	array := NewMapArray()

	_, err = array.Get(0)
	assert.ErrorIs(t, err, constants.ErrNotFound)
	assert.Equal(t, 0, array.Length())

	foo := Item{Value: "foo"}
	bar := Item{Value: "bar"}
	acme := Item{Value: "acme"}

	// PUSH
	array = array.Push(foo)
	assert.Equal(t, 1, array.Length())

	array = array.Push(bar)
	assert.Equal(t, 2, array.Length())

	array = array.Push(acme)
	assert.Equal(t, 3, array.Length())

	// GET
	item, err = array.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, foo, item)

	item, err = array.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, bar, item)

	item, err = array.Get(2)
	assert.NoError(t, err)
	assert.Equal(t, acme, item)

	item, err = array.Get(3)
	assert.ErrorIs(t, err, constants.ErrNotFound)

	// POP
	item, err = array.Pop()
	assert.NoError(t, err)
	assert.Equal(t, acme, item)

	item, err = array.Pop()
	assert.NoError(t, err)
	assert.Equal(t, bar, item)

	item, err = array.Pop()
	assert.NoError(t, err)
	assert.Equal(t, foo, item)

	item, err = array.Pop()
	assert.ErrorIs(t, err, constants.ErrEmpty)

	// DELETE AT
	array = array.Push(foo)
	array = array.Push(bar)
	array = array.Push(acme)

	item, err = array.DeleteAt(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, array.Length())

	item, err = array.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, foo, item)

	item, err = array.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, acme, item)

	item, err = array.Get(2)
	assert.ErrorIs(t, err, constants.ErrNotFound)

	item, err = array.DeleteAt(5)
	assert.ErrorIs(t, err, constants.ErrNotFound)
}
