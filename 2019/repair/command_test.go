package repair

import (
	"testing"

	"github.com/mfesenko/adventofcode/2019/math"
	"github.com/stretchr/testify/assert"
)

func TestApply(t *testing.T) {
	type testData struct {
		command MovementCommand
		dx      int64
		dy      int64
	}

	tests := []testData{
		{
			command: North,
			dy:      1,
		},
		{
			command: South,
			dy:      -1,
		},
		{
			command: West,
			dx:      -1,
		},
		{
			command: East,
			dx:      1,
		},
		{
			command: Invalid,
		},
	}

	position := math.NewPoint(12, 34)
	for _, test := range tests {
		newPosition := test.command.Apply(position)
		assert.Equal(t, newPosition.X, position.X+test.dx)
		assert.Equal(t, newPosition.Y, position.Y+test.dy)
	}
}

func TestReverse(t *testing.T) {
	testData := map[MovementCommand]MovementCommand{
		North:   South,
		South:   North,
		West:    East,
		East:    West,
		Invalid: Invalid,
	}
	for command, reverse := range testData {
		assert.Equal(t, reverse, command.Reverse())
	}
}
