package arrays

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestDynamicArrayAppend(t *testing.T) {
	var item Item
	var err error

	da := NewDynamicArray(5)

	for i := 0; i < 10; i++ {
		da.Append(Item{Value: strconv.Itoa(i + 1)})
	}

	item, err = da.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, "2", item.Value)

	item, err = da.Get(7)
	assert.NoError(t, err)
	assert.Equal(t, "8", item.Value)
}
