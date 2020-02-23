package navigation

// ShortestPathAlg encapsulates an algorithm for calculating the shortest path
type ShortestPathAlg struct {
	visited   map[string]bool
	distances map[string]int
}

// NewShortestPathAlg creates a ShortestPathAlg
func NewShortestPathAlg() *ShortestPathAlg {
	return &ShortestPathAlg{
		visited:   map[string]bool{},
		distances: map[string]int{},
	}
}

// GetDistance returns the distance to the dest node
func (s *ShortestPathAlg) GetDistance(dest *Node) int {
	return s.distances[dest.name]
}

// Run calculates shortest path from src node to all the other nodes in the map
func (s *ShortestPathAlg) Run(src *Node) {
	nodes := []*Node{src}
	for i := 0; i < len(nodes); i++ {
		nodes = append(nodes, s.processNode(nodes[i])...)
	}
}

func (s *ShortestPathAlg) getNeighbours(node *Node) []*Node {
	var nodes []*Node
	for _, leaf := range node.leaves {
		if !s.visited[leaf.name] {
			nodes = append(nodes, leaf)
		}
	}
	if node.parent != nil && !s.visited[node.parent.name] {
		nodes = append(nodes, node.parent)
	}
	return nodes
}

func (s *ShortestPathAlg) processNode(node *Node) []*Node {
	nodes := s.getNeighbours(node)
	s.visited[node.name] = true
	for _, neighbour := range nodes {
		s.analyzeDistance(node.name, neighbour.name)
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
