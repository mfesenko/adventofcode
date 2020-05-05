package ascii

import (
	"testing"

	"github.com/mfesenko/adventofcode/2019/math"
	"github.com/stretchr/testify/assert"
)

func TestDirectionFromRune(t *testing.T) {
	tests := map[rune]Direction{
		lookingUp:    Up,
		lookingDown:  Down,
		lookingLeft:  Left,
		lookingRight: Right,
		'a':          Invalid,
	}
	for r, expected := range tests {
		assert.Equal(t, expected, DirectionFromRune(r))
	}
}

func TestMoveForward(t *testing.T) {
	position := math.NewPoint(23, 56)
	tests := map[Direction]math.Point{
		Up:      math.NewPoint(position.X, position.Y-1),
		Right:   math.NewPoint(position.X+1, position.Y),
		Down:    math.NewPoint(position.X, position.Y+1),
		Left:    math.NewPoint(position.X-1, position.Y),
		Invalid: position,
	}
	for direction, expected := range tests {
		assert.Equal(t, expected, direction.MoveForward(position))
	}
}

func TestTurnRight(t *testing.T) {
	turns := map[Direction]Direction{
		Up:      Right,
		Right:   Down,
		Down:    Left,
		Left:    Up,
		Invalid: Invalid,
	}
	for from, expected := range turns {
		assert.Equal(t, expected, from.TurnRight())
	}
}

func TestTurnLeft(t *testing.T) {
	turns := map[Direction]Direction{
		Up:      Left,
		Left:    Down,
		Down:    Right,
		Right:   Up,
		Invalid: Invalid,
	}
	for from, expected := range turns {
		assert.Equal(t, expected, from.TurnLeft())
	}
}
