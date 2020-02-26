package asteroid

import (
	"github.com/mfesenko/adventofcode/2019/math"
)

const _asteroid = '#'

// Map represents an easteroid map
type Map struct {
	width  int
	height int
	data   [][]bool
}

// NewMap creates a map
func NewMap(input []string) Map {
	height := len(input)
	width := len(input[0])
	data := make([][]bool, height)
	for y, row := range input {
		data[y] = make([]bool, width)
		for x, cell := range row {
			data[y][x] = cell == _asteroid
		}
	}
	return Map{
		width:  width,
		height: height,
		data:   data,
	}
}

// Width returns the width of the map
func (m Map) Width() int {
	return m.width
}

// Height returns the height of the map
func (m Map) Height() int {
	return m.height
}

// Asteroids returns locations of asteroids on the map
func (m Map) Asteroids() []math.Point {
	var points []math.Point
	for y, row := range m.data {
		for x, point := range row {
			if point {
				points = append(points, math.NewPoint(int32(x), int32(y)))
			}
		}
	}
	return points
}
