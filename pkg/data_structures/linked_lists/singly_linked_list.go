package linked_lists

import (
	"ds_algo_course/pkg/constants"
	"strings"
)

type SinglyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewSinglyLinkedList(initialValue string) LinkedList {
	head := Node{
		Value: initialValue,
		Next:  nil,
	}

	l := &SinglyLinkedList{
		Head:   &head,
		Tail:   &head,
		Length: 1,
	}

	return l
}

func (l *SinglyLinkedList) Append(value string) {
	newNode := Node{
		Value: value,
	}

	l.Tail.Next = &newNode
	l.Tail = &newNode
	l.Length++
}

func (l *SinglyLinkedList) Prepend(value string) {
	newNode := Node{
		Value: value,
	}

	newNode.Next = l.Head
	l.Head = &newNode
	l.Length++
}

func (l *SinglyLinkedList) Insert(index int, value string) error {
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
	previousNode.Next = &newNode

	l.Length++

	return nil
}

func (l *SinglyLinkedList) Lookup(index int) (*Node, error) {
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

func (l *SinglyLinkedList) Remove(index int) error {
	previousNode, err := l.Lookup(index - 1)
	if err != nil {
		return err
	}

	unwantedNode := previousNode.Next
	if unwantedNode == nil {
		return constants.ErrIndexOutOfBounds
	}

	previousNode.Next = unwantedNode.Next

	l.Length--

	return nil
}

func (l *SinglyLinkedList) String() string {
	var toArray []string
	var currentNode *Node

	currentNode = l.Head
	for currentNode != nil {
		toArray = append(toArray, currentNode.Value)
		currentNode = currentNode.Next
	}

	return strings.Join(toArray, ", ")
}

func (l *SinglyLinkedList) Reverse() {
	if l.Length <= 1 {
		return
	}

	first := l.Head
	second := first.Next

	l.Tail = first

	for second != nil {
		third := second.Next
		second.Next = first
		first = second
		second = third
	}

	l.Head.Next = nil
	l.Head = first
}
