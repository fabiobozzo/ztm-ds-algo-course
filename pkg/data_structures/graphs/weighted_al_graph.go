package graphs

type WEdge struct {
	NodeTo string
	Weight int
}

type AdjacencyListWGraph struct {
	size          int
	adjacencyList map[string][]WEdge
}

func NewAdjacencyListWGraph() WGraph {
	return &AdjacencyListWGraph{
		adjacencyList: map[string][]WEdge{},
	}
}

func (a *AdjacencyListWGraph) AddVertex(vertex string) {
	a.adjacencyList[vertex] = []WEdge{}
	a.size++
}

func (a *AdjacencyListWGraph) AddEdge(vertex1, vertex2 string, weight int) {
	if _, ok := a.adjacencyList[vertex1]; ok {
		a.adjacencyList[vertex1] = append(a.adjacencyList[vertex1], WEdge{NodeTo: vertex2, Weight: weight})
	}

	if _, ok := a.adjacencyList[vertex2]; ok {
		a.adjacencyList[vertex2] = append(a.adjacencyList[vertex2], WEdge{NodeTo: vertex1, Weight: weight})
	}
}

func (a *AdjacencyListWGraph) Vertexes() []string {
	var vertexes []string

	for v := range a.adjacencyList {
		vertexes = append(vertexes, v)
	}

	return vertexes
}

func (a *AdjacencyListWGraph) GetNeighbors(vertex string) []WEdge {
	if edges, found := a.adjacencyList[vertex]; found {
		return edges
	}

	return []WEdge{}
}
