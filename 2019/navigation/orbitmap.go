package navigation

import (
	"strings"
)

type (
	// Node represents a node on the map
	Node struct {
		name   string
		parent *Node
		leaves map[string]*Node
	}

	// OrbitMap represents an orbit map
	OrbitMap struct {
		nodes map[string]*Node
	}
)

// NewNode creates a new Node
func NewNode(name string) *Node {
	return &Node{
		name:   name,
		leaves: map[string]*Node{},
	}
}

// AddLeaf adds a leaf node to the node
func (n *Node) AddLeaf(leaf *Node) {
	leaf.parent = n
	n.leaves[leaf.name] = leaf
}

// CheckSum calculates check sum for the node
func (n *Node) CheckSum(start int) int {
	sum := len(n.leaves) * start
	for _, leaf := range n.leaves {
		sum += leaf.CheckSum(start+1) + 1
	}
	return sum
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
	alg := NewShortestPathAlg()
	alg.Run(m.nodes[srcName])
	return alg.GetDistance(m.nodes[destName])
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
