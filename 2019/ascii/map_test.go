package ascii

import (
	"testing"

	"github.com/mfesenko/adventofcode/2019/math"
	"github.com/stretchr/testify/assert"
)

func TestMapIntersections(t *testing.T) {
	scaffoldMap := loadTestMap()

	intersections := scaffoldMap.Intersections()

	expected := []math.Point{
		math.NewPoint(2, 2),
		math.NewPoint(2, 4),
		math.NewPoint(6, 4),
		math.NewPoint(10, 4),
	}
	assert.ElementsMatch(t, expected, intersections)
}

func TestMapGeneratePath(t *testing.T) {
	scaffoldMap := loadTestMap()

	path := scaffoldMap.GeneratePath()

	expected := []string{
		"4", "R", "2", "R", "2", "R", "12", "R", "2", "R", "6", "R", "4", "R", "4", "R", "6",
	}
	assert.Equal(t, expected, path)
}

func loadTestMap() *ScaffoldMap {
	input := `..#..........
..#..........
#######...###
#.#...#...#.#
#############
..#...#...#..
..#####...^..

`
	loader := NewMapLoader()
	for _, r := range input {
		loader.ProcessRune(r)
	}
	return loader.ScaffoldMap()
}
