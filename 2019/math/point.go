package math

// Point represents a point in 2D space
type Point struct {
	X int32
	Y int32
}

// NewPoint creates a point with provided coordinates
func NewPoint(x, y int32) Point {
	return Point{
		X: x,
		Y: y,
	}
}

// ManhattanDistance calculates Manhattan distance between the points
func (p Point) ManhattanDistance(other Point) int32 {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y)
}
