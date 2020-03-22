package math

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPointCreatesAPoint(t *testing.T) {
	x := rand.Int63()
	y := rand.Int63()

	point := NewPoint(x, y)

	assert.Equal(t, x, point.X)
	assert.Equal(t, y, point.Y)
}

func TestManhattanDistanceReturnsManhattanDistanceBetweenThePoints(t *testing.T) {
	a := NewPoint(0, 0)
	b := NewPoint(-2, 20)

	assert.Equal(t, int64(22), a.ManhattanDistance(b))
}

func TestString(t *testing.T) {
	a := NewPoint(rand.Int63(), rand.Int63())

	assert.Equal(t, fmt.Sprintf("{X: %v, Y: %v}", a.X, a.Y), a.String())
}
