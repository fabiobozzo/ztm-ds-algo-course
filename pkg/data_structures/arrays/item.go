package arrays

type Item struct {
	Value string
}

func (i Item) String() string {
	return i.Value
}
