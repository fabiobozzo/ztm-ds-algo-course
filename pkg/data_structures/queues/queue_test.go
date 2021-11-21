package queues

import (
	"ds_algo_course/pkg/constants"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewLinkedListQueue()

	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")

	assert.Equal(t, "a, b, c", fmt.Sprint(queue))

	a, err := queue.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, "a", a)

	b, err := queue.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, "b", b)

	c, err := queue.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, "c", c)

	_, err = queue.Dequeue()
	assert.ErrorIs(t, err, constants.ErrEmpty)
}
