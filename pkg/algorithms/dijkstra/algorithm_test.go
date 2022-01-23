package dijkstra

import (
	"ds_algo_course/pkg/data_structures/graphs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateShortestPath1(t *testing.T) {
	graph := graphs.NewAdjacencyListWGraph()

	graph.AddVertex("a")
	graph.AddVertex("b")
	graph.AddVertex("c")
	graph.AddVertex("d")
	graph.AddVertex("e")

	graph.AddEdge("a", "b", 7)
	graph.AddEdge("a", "c", 3)
	graph.AddEdge("c", "b", 1)
	graph.AddEdge("b", "d", 2)
	graph.AddEdge("b", "e", 6)
	graph.AddEdge("c", "d", 2)
	graph.AddEdge("d", "e", 4)

	shortestDistance, shortestPath := CalculateShortestPath(graph, "a", "e")
	assert.Equal(t, 9, shortestDistance)
	assert.Equal(t, []string{"a", "c", "d", "e"}, shortestPath)
}

func TestCalculateShortestPath2(t *testing.T) {
	graph := graphs.NewAdjacencyListWGraph()

	graph.AddVertex("0")
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddVertex("3")
	graph.AddVertex("4")
	graph.AddVertex("5")
	graph.AddVertex("6")
	graph.AddVertex("7")
	graph.AddVertex("8")

	graph.AddEdge("0", "1", 4)
	graph.AddEdge("0", "6", 7)
	graph.AddEdge("1", "6", 11)
	graph.AddEdge("1", "2", 9)
	graph.AddEdge("1", "7", 20)
	graph.AddEdge("6", "7", 1)
	graph.AddEdge("2", "4", 2)
	graph.AddEdge("4", "7", 1)
	graph.AddEdge("2", "3", 6)
	graph.AddEdge("7", "8", 3)
	graph.AddEdge("4", "3", 10)
	graph.AddEdge("4", "8", 5)
	graph.AddEdge("4", "5", 15)
	graph.AddEdge("3", "5", 5)
	graph.AddEdge("8", "5", 12)

	shortestDistance, shortestPath := CalculateShortestPath(graph, "0", "5")
	assert.Equal(t, 22, shortestDistance)
	assert.Equal(t, []string{"0", "6", "7", "4", "2", "3", "5"}, shortestPath)

	shortestDistance, shortestPath = CalculateShortestPath(graph, "0", "1")
	assert.Equal(t, 4, shortestDistance)
	assert.Equal(t, []string{"0", "1"}, shortestPath)

	shortestDistance, shortestPath = CalculateShortestPath(graph, "6", "1")
	assert.Equal(t, 11, shortestDistance)
	assert.Equal(t, []string{"6", "1"}, shortestPath)
}
