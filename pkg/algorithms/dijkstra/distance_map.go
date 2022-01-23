package dijkstra

type distanceMap struct {
	vertexes map[string]vertexDistance
}

type vertexDistance struct {
	shortestDistanceFromSource int
	previousVertex             string
}

func newDistanceMap(vertexes []string, from string) *distanceMap {
	dmap := distanceMap{vertexes: map[string]vertexDistance{}}

	for _, v := range vertexes {
		vertexDistance := vertexDistance{
			shortestDistanceFromSource: maxInt,
			previousVertex:             "",
		}

		if v == from {
			vertexDistance.shortestDistanceFromSource = 0
		}

		dmap.vertexes[v] = vertexDistance
	}

	return &dmap
}

func (m *distanceMap) GetShortestDistanceTo(vertex string) int {
	if v, exists := m.vertexes[vertex]; exists {
		return v.shortestDistanceFromSource
	}

	return maxInt
}

func (m *distanceMap) GetPreviousVertexOf(vertex string) string {
	if v, exists := m.vertexes[vertex]; exists {
		return v.previousVertex
	}

	return ""
}

func (m *distanceMap) SetShortestDistanceTo(vertex, previousVertex string, distance int) {
	m.vertexes[vertex] = vertexDistance{
		shortestDistanceFromSource: distance,
		previousVertex:             previousVertex,
	}
}

func (m *distanceMap) GetShortestPath(from, to string) []string {
	var reversePath []string

	reversePath = append(reversePath, to)
	previousNode := m.GetPreviousVertexOf(to)

	for previousNode != "" {
		reversePath = append(reversePath, previousNode)
		previousNode = m.GetPreviousVertexOf(previousNode)
	}

	for i, j := 0, len(reversePath)-1; i < j; i, j = i+1, j-1 {
		reversePath[i], reversePath[j] = reversePath[j], reversePath[i]
	}

	return reversePath
}
