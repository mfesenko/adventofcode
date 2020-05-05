package ascii

import (
	"github.com/mfesenko/adventofcode/2019/math"
)

// Direction represents a direction in which a robot is moving
type Direction int64

const (
	// Invalid represents an invalide direction value
	Invalid = -1
	// Up represents up direction
	Up Direction = 0
	// Right represents right direction
	Right Direction = 1
	// Down represents down direction
	Down Direction = 2
	// Left represents left direction
	Left Direction = 3
)

const (
	lookingUp    = '^'
	lookingDown  = 'v'
	lookingLeft  = '<'
	lookingRight = '>'
)

// DirectionFromRune returns a Direction that corresponds to a rune
func DirectionFromRune(r rune) Direction {
	switch r {
	case lookingUp:
		return Up
	case lookingRight:
		return Right
	case lookingDown:
		return Down
	case lookingLeft:
		return Left
	}
	return Invalid
}

// MoveForward returns next position for given direction
func (d Direction) MoveForward(position math.Point) math.Point {
	switch d {
	case Up:
		return math.NewPoint(position.X, position.Y-1)
	case Down:
		return math.NewPoint(position.X, position.Y+1)
	case Right:
		return math.NewPoint(position.X+1, position.Y)
	case Left:
		return math.NewPoint(position.X-1, position.Y)
	}
	return position
}

// TurnRight returns a direction for a robot after a right turn
func (d Direction) TurnRight() Direction {
	if d == Invalid {
		return Invalid
	}
	return ((d + 1) % 4)
}

// TurnLeft returns a direction for a robot after a left turn
func (d Direction) TurnLeft() Direction {
	if d == Invalid {
		return Invalid
	}
	return ((d + 3) % 4)
}
