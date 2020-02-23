package math

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIntervalCreatesAnInterval(t *testing.T) {
	start := NewPoint(rand.Int31(), rand.Int31())
	end := NewPoint(rand.Int31(), rand.Int31())

	interval := NewInterval(start, end)

	assert.Equal(t, start, interval.start)
	assert.Equal(t, end, interval.end)
}

func TestWhenLinesDoNotIntersectThenFindIntersectionReturnsFalse(t *testing.T) {
	a := NewInterval(NewPoint(1, 1), NewPoint(5, 1))
	b := NewInterval(NewPoint(2, 3), NewPoint(3, 3))

	i, ok := a.FindIntersection(b)

	assert.False(t, ok)
	assert.Equal(t, NewPoint(0, 0), i)
}

func TestWhenLinesDoNotIntersectOursideIntervalsThenFindIntersectionReturnsFalse(t *testing.T) {
	a := NewInterval(NewPoint(1, 1), NewPoint(5, 1))
	b := NewInterval(NewPoint(3, 2), NewPoint(3, 3))

	i, ok := a.FindIntersection(b)

	assert.False(t, ok)
	assert.Equal(t, NewPoint(0, 0), i)
}

func TestWhenIntervalsIntersectThenFindIntersectionReturnsIntersectionPoint(t *testing.T) {
	a := NewInterval(NewPoint(1, 1), NewPoint(5, 1))
	b := NewInterval(NewPoint(3, 0), NewPoint(3, 3))

	i, ok := a.FindIntersection(b)

	assert.True(t, ok)
	assert.Equal(t, NewPoint(3, 1), i)
}

func TestDistanceTo(t *testing.T) {
	a := NewInterval(NewPoint(1, 1), NewPoint(5, 1))

	assert.Equal(t, int32(3), a.DistanceTo(NewPoint(4, 1)))
}

func TestLength(t *testing.T) {
	a := NewInterval(NewPoint(1, 1), NewPoint(5, 1))

	assert.Equal(t, int32(4), a.Length())
}
