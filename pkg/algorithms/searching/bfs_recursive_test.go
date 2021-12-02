package searching

import (
	"ds_algo_course/pkg/data_structures/queues"
	"ds_algo_course/pkg/data_structures/trees"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecursiveBFS(t *testing.T) {
	bst := trees.NewBinarySearchTree()

	bst.Insert(9)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(20)
	bst.Insert(170)
	bst.Insert(15)
	bst.Insert(1)

	root := bst.GetRoot()

	q := queues.NewLinkedListQueue()
	q.Enqueue(root)

	var traversedNodes []int

	assert.Equal(t, []int{9, 4, 20, 1, 6, 15, 170}, recursiveBfs(q, traversedNodes))
}
