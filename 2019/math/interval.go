package math

// Interval represents an interval between two points
type Interval struct {
	start Point
	end   Point
}

// NewInterval creates an interval with provided start and end
func NewInterval(start, end Point) Interval {
	return Interval{
		start: start,
		end:   end,
	}
}

// Contains returns true if the provided point is within the interval
func (i Interval) Contains(point Point) bool {
	return Between(i.start.X, i.end.X, point.X) && Between(i.start.Y, i.end.Y, point.Y)
}

// Length returns the lenght of the interval
func (i Interval) Length() int32 {
	return i.DistanceTo(i.end)
}

// DistanceTo returns the distance from the start of the interval to the provided point
func (i Interval) DistanceTo(point Point) int32 {
	return Abs(point.X-i.start.X) + Abs(point.Y-i.start.Y)
}

// FindIntersection returns an intersection point for two intervals if it exists
// If the intersection does not exist it returns false
func (i Interval) FindIntersection(other Interval) (Point, bool) {
	k3 := (i.start.X-i.end.X)*(other.start.Y-other.end.Y) - (i.start.Y-i.end.Y)*(other.start.X-other.end.X)
	if k3 == 0 {
		return NewPoint(0, 0), false
	}
	k2 := other.start.X*other.end.Y - other.start.Y*other.end.X
	k1 := i.start.X*i.end.Y - i.start.Y*i.end.X
	x := (k1*(other.start.X-other.end.X) - k2*(i.start.X-i.end.X)) / k3
	y := (k1*(other.start.Y-other.end.Y) - k2*(i.start.Y-i.end.Y)) / k3

	point := NewPoint(x, y)
	if !i.Contains(point) || !other.Contains(point) {
		return NewPoint(0, 0), false
	}
	return point, true
}
