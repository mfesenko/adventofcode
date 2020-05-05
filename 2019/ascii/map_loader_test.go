package ascii

import (
	"testing"

	"github.com/mfesenko/adventofcode/2019/math"
	"github.com/stretchr/testify/assert"
)

func TestMapLoaderLoadsAMapFromRunes(t *testing.T) {
	l := NewMapLoader()

	runes := []rune{
		empty, scaffold, newLine,
		lookingDown, empty, newLine,
		newLine,
	}
	for _, r := range runes {
		l.ProcessRune(r)
	}

	scaffoldMap := l.ScaffoldMap()
	assert.Equal(t, math.NewPoint(0, 1), scaffoldMap.Position())
	assert.Equal(t, Down, scaffoldMap.Direction())
	assert.False(t, scaffoldMap.Contains(math.NewPoint(0, 0)))
	assert.True(t, scaffoldMap.Contains(math.NewPoint(1, 0)))
	assert.True(t, scaffoldMap.Contains(math.NewPoint(0, 1)))
	assert.False(t, scaffoldMap.Contains(math.NewPoint(1, 1)))
}

func TestMapLoaderStopProcessingRunesAfterTwoNewLines(t *testing.T) {
	l := NewMapLoader()
	assert.False(t, l.Done())
	l.ProcessRune(scaffold)
	assert.False(t, l.Done())
	l.ProcessRune(newLine)
	assert.False(t, l.Done())
	l.ProcessRune(newLine)
	assert.True(t, l.Done())
}
