package math

import (
	"fmt"
)

// Point represents a point in 2D space
type Point struct {
	X int64
	Y int64
}

// NewPoint creates a point with provided coordinates
func NewPoint(x, y int64) Point {
	return Point{
		X: x,
		Y: y,
	}
}

// ManhattanDistance calculates Manhattan distance between the points
func (p Point) ManhattanDistance(other Point) int64 {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y)
}

// String returns string representation of a point
func (p Point) String() string {
	return fmt.Sprintf("{X: %v, Y: %v}", p.X, p.Y)
}
