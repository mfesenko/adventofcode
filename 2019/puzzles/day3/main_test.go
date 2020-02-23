package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testData struct {
	input             []string
	manhattanDistance int32
	stepCount         int32
}

func tests() []testData {
	return []testData{
		{
			input: []string{
				"R8,U5,L5,D3",
				"U7,R6,D4,L4",
			},
			manhattanDistance: 6,
			stepCount:         30,
		},
		{
			input: []string{
				"R75,D30,R83,U83,L12,D49,R71,U7,L72",
				"U62,R66,U55,R34,D71,R55,D58,R83",
			},
			manhattanDistance: 159,
			stepCount:         610,
		},
		{
			input: []string{
				"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
				"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			},
			manhattanDistance: 135,
			stepCount:         410,
		},
	}
}

func TestPartOne(t *testing.T) {
	for _, test := range tests() {
		wires, err := parseWires(test.input)
		require.NoError(t, err)
		assert.Equal(t, test.manhattanDistance, partOne(wires))
	}
}

func TestPartTwo(t *testing.T) {
	for _, test := range tests() {
		wires, err := parseWires(test.input)
		require.NoError(t, err)
		assert.Equal(t, test.stepCount, partTwo(wires))
	}
}
