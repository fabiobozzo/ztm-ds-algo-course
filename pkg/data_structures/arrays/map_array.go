package arrays

import (
	"ds_algo_course/pkg/constants"
)

type MapArray struct {
	length int
	array  map[int]Item
}

func NewMapArray() Array {
	return &MapArray{
		length: 0,
		array:  map[int]Item{},
	}
}

func (m *MapArray) Get(index int) (Item, error) {
	item, found := m.array[index]
	if !found {
		return Item{}, constants.ErrNotFound
	}

	return item, nil
}

func (m *MapArray) Push(item Item) Array {
	m.array[m.length] = item
	m.length++

	return m
}

func (m *MapArray) Pop() (Item, error) {
	if m.length <= 0 {
		return Item{}, constants.ErrEmpty
	}

	item, _ := m.array[m.length-1]
	delete(m.array, m.length-1)
	m.length--

	return item, nil
}

func (m *MapArray) DeleteAt(index int) (Item, error) {
	item, found := m.array[index]
	if !found {
		return Item{}, constants.ErrNotFound
	}

	if err := m.shiftItems(index); err != nil {
		return Item{}, err
	}

	return item, nil
}

func (m *MapArray) shiftItems(index int) error {
	if index < 0 || index >= m.length {
		return constants.ErrIndexOutOfBounds
	}

	for i := index; i < m.length-1; i++ {
		m.array[i] = m.array[i+1]
	}

	delete(m.array, m.length-1)
	m.length--

	return nil
}

func (m *MapArray) Length() int {
	return m.length
}

func (m *MapArray) Append(item Item) Array {
	return m.Push(item)
}
