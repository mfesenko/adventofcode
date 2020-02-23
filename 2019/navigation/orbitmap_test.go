package navigation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func withTestOrbitMap(test func(*OrbitMap)) {
	orbitMap := LoadOrbitMap([]string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	})
	test(orbitMap)
}

func TestCheckSum(t *testing.T) {
	withTestOrbitMap(func(orbitMap *OrbitMap) {
		assert.Equal(t, 42, orbitMap.CheckSum())
	})
}

func TestFindParent(t *testing.T) {
	withTestOrbitMap(func(orbitMap *OrbitMap) {
		assert.Equal(t, "J", orbitMap.FindParent("K"))
	})
}

func TestFindShortestPath(t *testing.T) {
	withTestOrbitMap(func(orbitMap *OrbitMap) {
		assert.Equal(t, 4, orbitMap.FindShortestPath("K", "I"))
		assert.Equal(t, 3, orbitMap.FindShortestPath("E", "L"))
	})
}
