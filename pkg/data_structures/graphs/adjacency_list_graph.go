package graphs

import (
	"fmt"
	"sort"
	"strings"
)

type AdjacencyListGraph struct {
	size          int
	adjacencyList map[string][]string
}

func NewAdjacencyListGraph() Graph {
	return &AdjacencyListGraph{
		adjacencyList: map[string][]string{},
	}
}

func (a *AdjacencyListGraph) AddVertex(vertex string) {
	a.adjacencyList[vertex] = []string{}
	a.size++
}

func (a *AdjacencyListGraph) AddEdge(vertex1, vertex2 string) {
	if _, ok := a.adjacencyList[vertex1]; ok {
		a.adjacencyList[vertex1] = append(a.adjacencyList[vertex1], vertex2)
		sort.Strings(a.adjacencyList[vertex1])
	}
	if _, ok := a.adjacencyList[vertex2]; ok {
		a.adjacencyList[vertex2] = append(a.adjacencyList[vertex2], vertex1)
		sort.Strings(a.adjacencyList[vertex2])
	}
}

func (a *AdjacencyListGraph) ShowConnections() string {
	var connections []string
	for v, conn := range a.adjacencyList {
		nodeConnections := strings.Join(conn, `", "`)
		connections = append(connections, fmt.Sprintf(`{"node":%s,"connections":["%s"]`, v, nodeConnections))
	}

	return fmt.Sprintf("[%s]", strings.Join(connections, ","))
}
