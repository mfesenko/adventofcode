package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenFailedToParseDirectionThenParsePathReturnsError(t *testing.T) {
	_, err := parsePath("X22")

	assert.Error(t, err)
}

func TestWhenFailedToParseOffsetThenParsePathReturnsError(t *testing.T) {
	_, err := parsePath("UX22")

	assert.Error(t, err)
}

func TestWhenInputIsValidThenParsePathReturnsPath(t *testing.T) {
	type testData struct {
		input string
		dx    int64
		dy    int64
	}
	tests := []testData{
		{
			input: "U12",
			dy:    12,
		},
		{
			input: "D2",
			dy:    -2,
		},
		{
			input: "R100",
			dx:    100,
		},
		{
			input: "L23",
			dx:    -23,
		},
	}
	for _, test := range tests {
		path, err := parsePath(test.input)

		assert.NoError(t, err)
		assert.Equal(t, test.dx, path.dx)
		assert.Equal(t, test.dy, path.dy)
	}
}
