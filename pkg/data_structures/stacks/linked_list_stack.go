package stacks

import (
	"ds_algo_course/pkg/constants"
	"strings"
)

type LinkedListStack struct {
	Top    *Node
	Bottom *Node
	length int
}

func NewLinkedListStack() Stack {
	return &LinkedListStack{}
}

func (s *LinkedListStack) Peek() (string, error) {
	if s.length <= 0 {
		return "", constants.ErrEmpty
	}

	return s.Top.Value, nil
}

func (s *LinkedListStack) Push(value string) {
	newNode := Node{
		Value: value,
	}

	if s.length <= 0 {
		s.Top = &newNode
		s.Bottom = &newNode
		s.length = 1

		return
	}

	newNode.Next = s.Top
	s.Top = &newNode
	s.length++
}

func (s *LinkedListStack) Pop() (string, error) {
	if s.length <= 0 {
		return "", constants.ErrEmpty
	}

	node := *s.Top

	if s.length == 1 {
		s.Top = nil
		s.Bottom = nil
		s.length = 0

		return node.Value, nil
	}

	s.Top = s.Top.Next
	s.length--

	return node.Value, nil
}

func (s *LinkedListStack) IsEmpty() bool {
	return s.length == 0
}

func (s *LinkedListStack) String() string {
	var toArray []string
	var currentNode *Node

	currentNode = s.Top
	for currentNode != nil {
		toArray = append(toArray, currentNode.Value)
		currentNode = currentNode.Next
	}

	return strings.Join(toArray, ", ")
}

func (s *LinkedListStack) Length() int {
	return s.length
}
