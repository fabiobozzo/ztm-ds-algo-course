package arrays

import (
	"ds_algo_course/pkg/constants"
	"sync"
)

const maximumCap = 100

type DynamicArray struct {
	array []Item // here a slice is used as a fixed array (not using its `append` method)
	len   int    // the actual memory occupied by the items in the array
	cap   int    // the maximum capacity of the fixed array
	lock  sync.Mutex
}

func NewDynamicArray(cap int) Array {
	if cap <= 0 || cap > maximumCap {
		cap = maximumCap
	}

	return &DynamicArray{
		array: make([]Item, cap, cap),
		cap:   cap,
		len:   0,
	}
}

func (d *DynamicArray) Append(item Item) Array {
	d.lock.Lock()
	defer d.lock.Unlock()

	// length is equal to capacity, which means the array needs to be expanded
	if d.len >= d.cap {
		// the new capacity is twice the actual length (could be different, that's just a choice)
		newCap := 2 * d.len

		// if the previous capacity is 0, the new capacity is 1
		if d.cap == 0 {
			newCap = 1
		}

		// create new fixed size array and copy data into it
		newArray := make([]Item, newCap, newCap)
		for k, v := range d.array {
			newArray[k] = v
		}

		// replace underlying fixed array with the new larger one
		d.array = newArray
		d.cap = newCap

	}

	// insert new element in the array
	d.array[d.len] = item
	d.len++

	return d
}

func (d *DynamicArray) Get(index int) (Item, error) {
	if d.len == 0 || index >= d.len {
		return Item{}, constants.ErrIndexOutOfBounds
	}

	return d.array[index], nil
}

func (d *DynamicArray) Length() int {
	return d.len
}

func (d *DynamicArray) Push(_ Item) Array {
	panic("implement me")
}

func (d *DynamicArray) Pop() (Item, error) {
	panic("implement me")
}

func (d *DynamicArray) DeleteAt(_ int) (Item, error) {
	panic("implement me")
}
