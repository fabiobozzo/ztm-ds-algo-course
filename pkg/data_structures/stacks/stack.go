package stacks

type Node struct {
	Value string
	Next  *Node
}

type Stack interface {
	Peek() (string, error)
	Push(string)
	Pop() (string, error)
	IsEmpty() bool
	Length() int
}
