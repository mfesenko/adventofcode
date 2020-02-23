package main

import (
	"testing"

	"github.com/mfesenko/adventofcode/2019/math"
	"github.com/stretchr/testify/assert"
)

func TestWhenGivenInvalidInputThenParseWireReturnsAnError(t *testing.T) {
	wire, err := parseWire("asdf")

	assert.Error(t, err)
	assert.Nil(t, wire)
}

func TestWhenGivenValidInputThenParseWireReturnsAWire(t *testing.T) {
	expectedWire := wire{
		math.NewInterval(math.NewPoint(0, 0), math.NewPoint(8, 0)),
		math.NewInterval(math.NewPoint(8, 0), math.NewPoint(8, 5)),
		math.NewInterval(math.NewPoint(8, 5), math.NewPoint(3, 5)),
		math.NewInterval(math.NewPoint(3, 5), math.NewPoint(3, 2)),
	}

	wire, err := parseWire("R8,U5,L5,D3")

	assert.NoError(t, err)
	assert.Equal(t, expectedWire, wire)
}

func TestWhenGivenInvalidInputThenParseWiresReturnsAnError(t *testing.T) {
	wires, err := parseWires([]string{"R8,U5,L5,D3", "asdf"})

	assert.Error(t, err)
	assert.Nil(t, wires)
}

func TestWhenGivenValidInputThenParseWiresReturnsWires(t *testing.T) {
	input := []string{"R8,U5,L5,D3", "U62,R66,U55,R34,D71,R55,D58,R83"}

	wires, err := parseWires(input)

	assert.NoError(t, err)
	assert.Equal(t, len(input), len(wires))
}
