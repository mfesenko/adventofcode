package navigation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddLeafConnectsTwoNodes(t *testing.T) {
	a := NewNode("a")
	b := NewNode("b")

	require.Empty(t, a.leaves)
	require.Nil(t, b.parent)

	a.AddLeaf(b)

	assert.Equal(t, a, b.parent)
	assert.Equal(t, b, a.leaves[b.name])
}

func TestCheckSumForNodeWithNoLeavesReturnsZero(t *testing.T) {
	a := NewNode("a")

	assert.Equal(t, 0, a.CheckSum(2))
}

func TestCheckSumForNodeWithLeaves(t *testing.T) {
	a := NewNode("a")
	a.AddLeaf(NewNode("b"))
	a.AddLeaf(NewNode("c"))

	assert.Equal(t, 6, a.CheckSum(2))
}

func TestNeighboursReturnsEmptyResultForNodeWithoutLeavesAndParent(t *testing.T) {
	a := NewNode("a")
	assert.Empty(t, a.Neighbours())
}

func TestNeighboursReturnsLeavesForNodeWithLeavesAndWithoutParent(t *testing.T) {
	a := NewNode("a")
	a.AddLeaf(NewNode("b"))
	a.AddLeaf(NewNode("c"))

	assert.Equal(t, []string{"b", "c"}, a.Neighbours())
}

func TestNeighboursReturnsParentForNodeWithoutLeavesAndWithParent(t *testing.T) {
	a := NewNode("a")
	NewNode("b").AddLeaf(a)

	assert.Equal(t, []string{"b"}, a.Neighbours())
}

func TestNeighboursReturnsLeavesAndParentForNodeWithLeavesAndWithParent(t *testing.T) {
	a := NewNode("a")
	NewNode("b").AddLeaf(a)
	a.AddLeaf(NewNode("c"))
	a.AddLeaf(NewNode("d"))

	assert.Equal(t, []string{"b", "c", "d"}, a.Neighbours())
}
