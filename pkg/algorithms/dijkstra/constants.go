package dijkstra

const (
	uintSize = 32 << (^uint(0) >> 32 & 1)
	maxInt   = 1<<(uintSize-1) - 1
)
