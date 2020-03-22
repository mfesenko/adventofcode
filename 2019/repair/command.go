package repair

import (
	"github.com/mfesenko/adventofcode/2019/math"
)

// MovementCommand represents a command
type MovementCommand int64

const (
	// Invalid represents an invalid movement command
	Invalid = MovementCommand(0)
	// North represents a command for moving north
	North = MovementCommand(1)
	// South represents a command for moving south
	South = MovementCommand(2)
	// West represents a command for moving west
	West = MovementCommand(3)
	// East represents a command for moving east
	East = MovementCommand(4)
)

// Apply returns a new position after applying a movement command to given position
func (c MovementCommand) Apply(position math.Point) math.Point {
	switch c {
	case North:
		return math.NewPoint(position.X, position.Y+1)
	case South:
		return math.NewPoint(position.X, position.Y-1)
	case East:
		return math.NewPoint(position.X+1, position.Y)
	case West:
		return math.NewPoint(position.X-1, position.Y)
	}
	return position
}

// Reverse returns a corresponding opposite movement command
func (c MovementCommand) Reverse() MovementCommand {
	switch c {
	case North:
		return South
	case South:
		return North
	case West:
		return East
	case East:
		return West
	}
	return Invalid
}
