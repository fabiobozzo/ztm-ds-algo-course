package stacks

import (
	"ds_algo_course/pkg/constants"
	"strings"
)

type ArrayStack struct {
	array  []string
	length int
}

func NewArrayStack(cap int) Stack {
	return &ArrayStack{
		array: make([]string, cap),
	}
}

func (a *ArrayStack) Peek() (string, error) {
	if a.length <= 0 {
		return "", constants.ErrEmpty
	}

	return a.array[a.length-1], nil
}

func (a *ArrayStack) Push(s string) {
	a.array[a.length] = s
	a.length++
}

func (a *ArrayStack) Pop() (string, error) {
	if a.length <= 0 {
		return "", constants.ErrEmpty
	}

	element := a.array[a.length-1]
	a.array[a.length-1] = ""
	a.length--

	return element, nil
}

func (a *ArrayStack) IsEmpty() bool {
	return a.length <= 0
}

func (a *ArrayStack) Length() int {
	return a.length
}

func (a *ArrayStack) String() string {
	return strings.Join(a.array, ", ")
}
