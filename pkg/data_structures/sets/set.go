package sets

type Set interface {
	Put(string)
	Remove(string)
	Exists(string) bool
	Items() []string
	Size() int
}
