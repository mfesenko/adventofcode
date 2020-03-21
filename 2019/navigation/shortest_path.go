package navigation

type (
	// Graph represents a graph
	Graph interface {
		Neighbours(string) []string
	}

	// ShortestPathAlg encapsulates an algorithm for calculating the shortest path
	ShortestPathAlg struct {
		graph     Graph
		visited   map[string]bool
		distances map[string]int
	}
)

// NewShortestPathAlg creates a ShortestPathAlg
func NewShortestPathAlg(graph Graph) *ShortestPathAlg {
	return &ShortestPathAlg{
		graph:     graph,
		visited:   map[string]bool{},
		distances: map[string]int{},
	}
}

// GetDistance returns the distance to the dest node
func (s *ShortestPathAlg) GetDistance(dest string) int {
	return s.distances[dest]
}

// Run calculates shortest path from src node to all the other nodes in the map
func (s *ShortestPathAlg) Run(src string) {
	nodes := []string{src}
	for i := 0; i < len(nodes); i++ {
		nodes = append(nodes, s.processNode(nodes[i])...)
	}
}

func (s *ShortestPathAlg) processNode(node string) []string {
	s.visited[node] = true
	neighbours := s.graph.Neighbours(node)
	var nodes []string
	for _, neighbour := range neighbours {
		if s.visited[neighbour] {
			continue
		}
		s.analyzeDistance(node, neighbour)
		nodes = append(nodes, neighbour)
	}
	return nodes
}

func (s *ShortestPathAlg) analyzeDistance(fromNode string, node string) {
	curDistance := s.distances[fromNode] + 1
	distance, ok := s.distances[node]
	if !ok || curDistance < distance {
		s.distances[node] = curDistance
	}
}
