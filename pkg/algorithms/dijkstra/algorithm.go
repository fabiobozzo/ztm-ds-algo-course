package dijkstra

import (
	"ds_algo_course/pkg/data_structures/graphs"
	"ds_algo_course/pkg/data_structures/sets"
)

func CalculateShortestPath(graph graphs.WGraph, from, to string) (int, []string) {
	vertexes := graph.Vertexes()
	unvisited := sets.NewHashSetWith(vertexes)

	distanceMap := newDistanceMap(vertexes, from)

	for unvisited.Size() > 1 {
		currentNode := findNodeWithMinimumDistance(unvisited, distanceMap)
		currentNodeDistance := distanceMap.GetShortestDistanceTo(currentNode)

		for _, neighbor := range graph.GetNeighbors(currentNode) {
			if unvisited.Exists(neighbor.NodeTo) {
				currentDistanceToNeighbor := distanceMap.GetShortestDistanceTo(neighbor.NodeTo)
				distanceToNeighbor := currentNodeDistance + neighbor.Weight

				if distanceToNeighbor < currentDistanceToNeighbor {
					distanceMap.SetShortestDistanceTo(neighbor.NodeTo, currentNode, distanceToNeighbor)
				}
			}
		}

		unvisited.Remove(currentNode)
	}

	unvisited.Remove(to)

	return distanceMap.GetShortestDistanceTo(to), distanceMap.GetShortestPath(from, to)
}

func findNodeWithMinimumDistance(unvisited sets.Set, distanceMap *distanceMap) string {
	var node string

	if unvisited.Size() == 0 {
		return ""
	}

	if unvisited.Size() == 1 {
		return unvisited.Items()[0]
	}

	minShortestDistance := maxInt

	for _, v := range unvisited.Items() {
		distanceFromSource := distanceMap.GetShortestDistanceTo(v)

		if distanceFromSource < minShortestDistance {
			minShortestDistance = distanceFromSource
			node = v
		}
	}

	return node
}
