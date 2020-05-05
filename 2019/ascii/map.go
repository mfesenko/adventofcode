package ascii

import (
	"strconv"

	"github.com/mfesenko/adventofcode/2019/math"
)

// ScaffoldMap represents a view of the scaffolds visible to a robot
type ScaffoldMap struct {
	scaffolds map[math.Point]struct{}
	position  math.Point
	direction Direction
}

// NewScaffoldMap creates a ScaffoldMap
func NewScaffoldMap() *ScaffoldMap {
	return &ScaffoldMap{
		scaffolds: map[math.Point]struct{}{},
	}
}

// SetPosition updates position of the robot on the map
func (m *ScaffoldMap) SetPosition(position math.Point) {
	m.position = position
}

// Position returns position of the robot on the map
func (m *ScaffoldMap) Position() math.Point {
	return m.position
}

// SetDirection update direction the robot is facing
func (m *ScaffoldMap) SetDirection(direction Direction) {
	m.direction = direction
}

// Direction returns direction the robot is facing
func (m *ScaffoldMap) Direction() Direction {
	return m.direction
}

// Contains returns true if there is a scaffold on the given position
func (m *ScaffoldMap) Contains(position math.Point) bool {
	_, ok := m.scaffolds[position]
	return ok
}

// AddScaffold adds a scaffold to the map
func (m *ScaffoldMap) AddScaffold(position math.Point) {
	m.scaffolds[position] = struct{}{}
}

// Intersections returns all intercestions of the scaffolds in the map
func (m *ScaffoldMap) Intersections() []math.Point {
	var intersections []math.Point
	for s := range m.scaffolds {
		neighbours := []math.Point{
			math.NewPoint(s.X-1, s.Y),
			math.NewPoint(s.X+1, s.Y),
			math.NewPoint(s.X, s.Y-1),
			math.NewPoint(s.X, s.Y+1),
		}
		count := 0
		for _, n := range neighbours {
			if _, ok := m.scaffolds[n]; ok {
				count++
			}
		}
		if count > 2 {
			intersections = append(intersections, s)
		}
	}
	return intersections
}

// GeneratePath returns a path that a robot needs to take in order to visit every scaffold at least once
func (m *ScaffoldMap) GeneratePath() []string {
	position := m.position
	direction := m.direction
	path := make([]string, 0)
	stepCount := 0
	for {
		forwardPosition := direction.MoveForward(position)
		if m.Contains(forwardPosition) {
			stepCount++
			position = forwardPosition
			continue
		}

		rightPosition := direction.TurnRight().MoveForward(position)
		if m.Contains(rightPosition) {
			direction = direction.TurnRight()
			if stepCount > 0 {
				path = append(path, strconv.Itoa(stepCount))
			}
			path = append(path, "R")
			stepCount = 0
			continue
		}

		leftPosition := direction.TurnLeft().MoveForward(position)
		if m.Contains(leftPosition) {
			direction = direction.TurnLeft()
			if stepCount > 0 {
				path = append(path, strconv.Itoa(stepCount))
			}
			path = append(path, "L")
			stepCount = 0
			continue
		}

		if stepCount > 0 {
			path = append(path, strconv.Itoa(stepCount))
		}
		break
	}
	return path
}
