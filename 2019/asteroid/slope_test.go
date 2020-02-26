package asteroid

import (
	"testing"

	"github.com/mfesenko/adventofcode/2019/math"
	"github.com/stretchr/testify/assert"
)

func TestCalculateSlope(t *testing.T) {
	expectedSlope := slope{dx: 7, dy: -5}
	slope := calculateSlope(math.NewPoint(21, -15), math.NewPoint(0, 0))
	assert.Equal(t, expectedSlope, slope)
}
