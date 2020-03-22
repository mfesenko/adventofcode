package repair

import (
	"math/rand"
	"testing"

	"github.com/mfesenko/adventofcode/2019/math"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func withAreaMap(test func(*areaMap, math.Point)) {
	areaMap := newAreaMap()
	position := math.NewPoint(rand.Int63(), rand.Int63())
	test(areaMap, position)
}

func TestWhenAreaMapDoesNotContainAPositionThenContainsReturnsFalse(t *testing.T) {
	withAreaMap(func(areaMap *areaMap, position math.Point) {
		assert.False(t, areaMap.Contains(position))
	})
}
func TestWhenAreaMapContainsAPositionThenContainsReturnsTrue(t *testing.T) {
	withAreaMap(func(areaMap *areaMap, position math.Point) {
		areaMap.Update(position, HitWall)
		assert.True(t, areaMap.Contains(position))
	})
}

func TestWhenAreaMapDoesNotContainAPositionThenNeighboursReturnsNil(t *testing.T) {
	withAreaMap(func(areaMap *areaMap, position math.Point) {
		assert.Nil(t, areaMap.Neighbours(position.String()))
	})
}

func TestWhenAreaMapContainsAPositionThenNeighboursReturnsOnlyExistingNotWallNeighbours(t *testing.T) {
	withAreaMap(func(areaMap *areaMap, position math.Point) {
		areaMap.Update(position, MovedOneStep)
		knownNeighbours := map[StatusCode]math.Point{
			HitWall:           math.NewPoint(position.X, position.Y-1),
			MovedOneStep:      math.NewPoint(position.X, position.Y+1),
			FoundOxygenSystem: math.NewPoint(position.X-1, position.Y),
		}
		for status, neighbour := range knownNeighbours {
			areaMap.Update(neighbour, status)
		}

		neighbours := areaMap.Neighbours(position.String())
		require.Equal(t, 2, len(neighbours))
		assert.Contains(t, neighbours, knownNeighbours[MovedOneStep].String())
		assert.Contains(t, neighbours, knownNeighbours[FoundOxygenSystem].String())
	})
}
