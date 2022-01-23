package sets

type HashSet struct {
	items map[string]bool
}

func NewHashSet() Set {
	return &HashSet{items: map[string]bool{}}
}

func NewHashSetWith(items []string) Set {
	set := NewHashSet()

	for _, i := range items {
		set.Put(i)
	}

	return set
}

func (s *HashSet) Put(i string) {
	s.items[i] = true
}

func (s *HashSet) Remove(i string) {
	delete(s.items, i)
}

func (s *HashSet) Exists(i string) bool {
	_, exists := s.items[i]

	return exists
}

func (s *HashSet) Items() []string {
	var items []string

	for i := range s.items {
		items = append(items, i)
	}

	return items
}

func (s *HashSet) Size() int {
	return len(s.items)
}
