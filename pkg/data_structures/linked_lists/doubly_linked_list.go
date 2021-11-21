package linked_lists

import (
	"ds_algo_course/pkg/constants"
	"strings"
)

type DoublyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewDoublyLinkedList(initialValue string) LinkedList {
	head := Node{
		Value:    initialValue,
		Next:     nil,
		Previous: nil,
	}

	l := &DoublyLinkedList{
		Head:   &head,
		Tail:   &head,
		Length: 1,
	}

	return l
}

func (l *DoublyLinkedList) Append(value string) {
	newNode := Node{
		Value: value,
	}

	newNode.Previous = l.Tail
	l.Tail.Next = &newNode
	l.Tail = &newNode
	l.Length++
}

func (l *DoublyLinkedList) Prepend(value string) {
	newNode := Node{
		Value: value,
	}

	newNode.Next = l.Head
	l.Head.Previous = &newNode
	l.Head = &newNode
	l.Length++
}

func (l *DoublyLinkedList) Insert(index int, value string) error {
	if index <= 0 {
		l.Prepend(value)

		return nil
	}

	if index >= l.Length {
		l.Append(value)

		return nil
	}

	previousNode, err := l.Lookup(index - 1)
	if err != nil {
		return err
	}
	nextNode := previousNode.Next

	newNode := Node{
		Value: value,
		Next:  nextNode,
	}
	newNode.Previous = previousNode
	nextNode.Previous = &newNode
	previousNode.Next = &newNode

	l.Length++

	return nil
}

func (l *DoublyLinkedList) Lookup(index int) (*Node, error) {
	if index < 0 || index >= l.Length {
		return nil, constants.ErrIndexOutOfBounds
	}

	i := 0
	currentNode := l.Head
	for i != index {
		currentNode = currentNode.Next
		i++
	}

	return currentNode, nil
}

func (l *DoublyLinkedList) Remove(index int) error {
	previousNode, err := l.Lookup(index - 1)
	if err != nil {
		return err
	}

	unwantedNode := previousNode.Next
	if unwantedNode == nil {
		return constants.ErrIndexOutOfBounds
	}

	previousNode.Next = unwantedNode.Next
	unwantedNode.Next.Previous = previousNode

	l.Length--

	return nil
}

func (l *DoublyLinkedList) String() string {
	var toArray []string
	var currentNode *Node

	currentNode = l.Head
	for currentNode != nil {
		toArray = append(toArray, currentNode.Value)
		currentNode = currentNode.Next
	}

	return strings.Join(toArray, ", ")
}

func (l *DoublyLinkedList) Reverse() {
	panic("implement me")
}
