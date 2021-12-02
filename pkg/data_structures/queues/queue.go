package queues

import (
	"ds_algo_course/pkg/constants"
	"strings"
)

type Queue interface {
	Peek() (interface{}, error)
	Enqueue(interface{})
	Dequeue() (interface{}, error)
	Length() int
}

type Node struct {
	Value interface{}
	Next  *Node
}

type linkedListQueue struct {
	first  *Node
	last   *Node
	length int
}

func NewLinkedListQueue() Queue {
	return &linkedListQueue{}
}

func (l *linkedListQueue) Peek() (interface{}, error) {
	if l.length <= 0 {
		return "", constants.ErrEmpty
	}

	return l.first.Value, nil
}

func (l *linkedListQueue) Enqueue(value interface{}) {
	node := Node{
		Value: value,
	}

	if l.length <= 0 {
		l.first = &node
		l.last = &node
		l.length = 1

		return
	}

	l.last.Next = &node
	l.last = &node
	l.length++
}

func (l *linkedListQueue) Dequeue() (interface{}, error) {
	if l.length <= 0 {
		return "", constants.ErrEmpty
	}

	if l.length == 1 {
		l.last = nil
	}

	dequeued := *l.first
	l.first = l.first.Next
	l.length--

	return dequeued.Value, nil
}

func (l *linkedListQueue) Length() int {
	return l.length
}

func (l *linkedListQueue) String() string {
	var toArray []string
	var currentNode *Node

	currentNode = l.first
	for currentNode != nil {
		toArray = append(toArray, currentNode.Value.(string))
		currentNode = currentNode.Next
	}

	return strings.Join(toArray, ", ")
}
