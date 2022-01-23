package graphs

type WGraph interface {
	AddVertex(vertex string)
	AddEdge(vertex1, vertex2 string, weight int)
	GetNeighbors(vertex string) []WEdge
	Vertexes() []string
}
