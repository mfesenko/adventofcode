package navigation

import (
	"strings"
)

// OrbitMap represents an orbit map
type OrbitMap struct {
	nodes map[string]*Node
}

func (m *OrbitMap) root() *Node {
	return m.nodes["COM"]
}

// CheckSum calculates the check sum for the map
func (m *OrbitMap) CheckSum() int {
	return m.root().CheckSum(0)
}

// FindParent finds the name of the parent for the node
func (m *OrbitMap) FindParent(name string) string {
	return m.nodes[name].parent.name
}

// FindShortestPath finds shortest path between two nodes
func (m *OrbitMap) FindShortestPath(srcName string, destName string) int {
	alg := NewShortestPathAlg(m)
	alg.Run(srcName)
	return alg.GetDistance(destName)
}

// Neighbours returns neighbours for the given node
func (m *OrbitMap) Neighbours(nodeName string) []string {
	if node, ok := m.nodes[nodeName]; ok {
		return node.Neighbours()
	}
	return nil
}

// LoadOrbitMap loads orbit map from input data
func LoadOrbitMap(data []string) *OrbitMap {
	nodes := map[string]*Node{}
	for _, line := range data {
		currentNodes := strings.Split(line, ")")
		center := findOrCreateNode(nodes, currentNodes[0])
		orbiting := findOrCreateNode(nodes, currentNodes[1])
		center.AddLeaf(orbiting)
	}

	return &OrbitMap{
		nodes: nodes,
	}
}

func findOrCreateNode(nodes map[string]*Node, nodeName string) *Node {
	node, ok := nodes[nodeName]
	if ok {
		return node
	}

	node = NewNode(nodeName)
	nodes[nodeName] = node
	return node
}
