package math

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPointCreatesAPoint(t *testing.T) {
	x := rand.Int31()
	y := rand.Int31()

	point := NewPoint(x, y)

	assert.Equal(t, x, point.X)
	assert.Equal(t, y, point.Y)
}

func TestManhattanDistanceReturnsManhattanDistanceBetweenThePoints(t *testing.T) {
	a := NewPoint(0, 0)
	b := NewPoint(-2, 20)

	assert.Equal(t, int32(22), a.ManhattanDistance(b))
}
