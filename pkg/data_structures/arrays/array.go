package arrays

type Array interface {
	Get(index int) (Item, error)
	Push(item Item) Array
	Pop() (Item, error)
	DeleteAt(index int) (Item, error)
	Length() int
	Append(item Item) Array
}
