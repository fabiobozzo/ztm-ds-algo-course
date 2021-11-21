package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	bst := NewBinarySearchTree()

	bst.Insert(6)
	bst.Insert(8)
	bst.Insert(7)
	bst.Insert(10)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(1)
	bst.Insert(9)

	root := bst.GetRoot()

	assert.Equal(t, 6, root.Value)
	assert.Equal(t, 8, root.Right.Value)
	assert.Equal(t, 7, root.Right.Left.Value)
	assert.Equal(t, 10, root.Right.Right.Value)
	assert.Equal(t, 2, root.Left.Value)
	assert.Equal(t, 4, root.Left.Right.Value)
	assert.Equal(t, 1, root.Left.Left.Value)
	assert.Equal(t, 9, root.Right.Right.Left.Value)
}

func TestLookup(t *testing.T) {
	bst := NewBinarySearchTree()

	bst.Insert(6)
	bst.Insert(8)
	bst.Insert(7)
	bst.Insert(10)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(1)
	bst.Insert(9)

	assert.Equal(t, 6, bst.Lookup(6).Value)
	assert.Equal(t, 8, bst.Lookup(8).Value)
	assert.Equal(t, 7, bst.Lookup(7).Value)
	assert.Equal(t, 10, bst.Lookup(10).Value)
	assert.Equal(t, 2, bst.Lookup(2).Value)
	assert.Equal(t, 4, bst.Lookup(4).Value)
	assert.Equal(t, 1, bst.Lookup(1).Value)
	assert.Equal(t, 9, bst.Lookup(9).Value)

	assert.Nil(t, bst.Lookup(11))
}
