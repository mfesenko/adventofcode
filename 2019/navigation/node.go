package navigation

import (
	"sort"
)

// Node represents a node on the map
type Node struct {
	name   string
	parent *Node
	leaves map[string]*Node
}

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

// Neighbours returns neighbours for the node
func (n *Node) Neighbours() []string {
	var neighbours []string
	for name := range n.leaves {
		neighbours = append(neighbours, name)
	}
	if n.parent != nil {
		neighbours = append(neighbours, n.parent.name)
	}
	sort.Strings(neighbours)
	return neighbours
}
