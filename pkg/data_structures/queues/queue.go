package queues

import (
	"ds_algo_course/pkg/constants"
	"strings"
)

type Queue interface {
	Peek() (string, error)
	Enqueue(string)
	Dequeue() (string, error)
}

type Node struct {
	Value string
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

func (l *linkedListQueue) Peek() (string, error) {
	if l.length <= 0 {
		return "", constants.ErrEmpty
	}

	return l.first.Value, nil
}

func (l *linkedListQueue) Enqueue(value string) {
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

func (l *linkedListQueue) Dequeue() (string, error) {
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

func (l *linkedListQueue) String() string {
	var toArray []string
	var currentNode *Node

	currentNode = l.first
	for currentNode != nil {
		toArray = append(toArray, currentNode.Value)
		currentNode = currentNode.Next
	}

	return strings.Join(toArray, ", ")
}
