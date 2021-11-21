package graphs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraph(t *testing.T) {
	graph := NewAdjacencyListGraph()

	graph.AddVertex("0")
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddVertex("3")
	graph.AddVertex("4")
	graph.AddVertex("5")
	graph.AddVertex("6")
	graph.AddEdge("3", "1")
	graph.AddEdge("3", "4")
	graph.AddEdge("4", "2")
	graph.AddEdge("4", "5")
	graph.AddEdge("1", "2")
	graph.AddEdge("1", "0")
	graph.AddEdge("0", "2")
	graph.AddEdge("6", "5")

	connections := graph.ShowConnections()
	assert.Equal(t, `[{"node":0,"connections":["1", "2"],{"node":1,"connections":["0", "2", "3"],{"node":2,"connections":["0", "1", "4"],{"node":3,"connections":["1", "4"],{"node":4,"connections":["2", "3", "5"],{"node":5,"connections":["4", "6"],{"node":6,"connections":["5"]]`, connections)
}
