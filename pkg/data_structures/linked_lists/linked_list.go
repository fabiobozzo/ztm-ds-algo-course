package linked_lists

type LinkedList interface {
	Append(value string)
	Prepend(value string)
	Insert(index int, value string) error
	Lookup(index int) (*Node, error)
	Remove(index int) error
	Reverse()
}

type Node struct {
	Value    string
	Next     *Node
	Previous *Node // only for doubly LL
}
