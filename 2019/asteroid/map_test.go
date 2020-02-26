package asteroid

import (
	"testing"

	"github.com/mfesenko/adventofcode/2019/math"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	m := NewMap([]string{
		"#..",
		"...",
		".#.",
		"##.",
	})

	assert.Equal(t, 4, m.Height())
	assert.Equal(t, 3, m.Width())
	assert.Equal(t, []math.Point{
		math.NewPoint(0, 0),
		math.NewPoint(1, 2),
		math.NewPoint(0, 3),
		math.NewPoint(1, 3),
	}, m.Asteroids())
}
