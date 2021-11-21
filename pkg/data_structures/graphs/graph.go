package graphs

type Graph interface {
	AddVertex(vertex string)
	AddEdge(vertex1, vertex2 string)
	ShowConnections() string
}
