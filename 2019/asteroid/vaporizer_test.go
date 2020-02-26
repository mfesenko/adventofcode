package asteroid

import (
	"testing"

	"github.com/mfesenko/adventofcode/2019/math"
	"github.com/stretchr/testify/assert"
)

func TestVaporize(t *testing.T) {
	vaporizer := NewVaporizer(NewMap([]string{
		".#....#####...#..",
		"##...##.#####..##",
		"##...#...#.#####.",
		"..#.....#...###..",
		"..#.#.....#....##",
	}))
	location := math.NewPoint(8, 3)
	asteroids := vaporizer.Vaporize(location)
	expectedOrder := []math.Point{
		math.NewPoint(8, 1),  // .#....###24...#..
		math.NewPoint(9, 0),  // ##...##.13#67..9#
		math.NewPoint(9, 1),  // ##...#...5.8####.
		math.NewPoint(10, 0), // ..#.....X...###..
		math.NewPoint(9, 2),  // ..#.#.....#....##
		math.NewPoint(11, 1),
		math.NewPoint(12, 1),
		math.NewPoint(11, 2),
		math.NewPoint(15, 1),

		math.NewPoint(12, 2), // .#....###.....#..
		math.NewPoint(13, 2), // ##...##...#.....#
		math.NewPoint(14, 2), // ##...#......1234.
		math.NewPoint(15, 2), // ..#.....X...5##..
		math.NewPoint(12, 3), // ..#.9.....8....76
		math.NewPoint(16, 4),
		math.NewPoint(15, 4),
		math.NewPoint(10, 4),
		math.NewPoint(4, 4),

		math.NewPoint(2, 4), // .8....###.....#..
		math.NewPoint(2, 3), // 56...9#...#.....#
		math.NewPoint(0, 2), // 34...7...........
		math.NewPoint(1, 2), // ..2.....X....##..
		math.NewPoint(0, 1), // ..1..............
		math.NewPoint(1, 1),
		math.NewPoint(5, 2),
		math.NewPoint(1, 0),
		math.NewPoint(5, 1),

		math.NewPoint(6, 1),  // ......234.....6..
		math.NewPoint(6, 0),  // ......1...5.....7
		math.NewPoint(7, 0),  // .................
		math.NewPoint(8, 0),  // ........X....89..
		math.NewPoint(10, 1), // .................
		math.NewPoint(14, 0),
		math.NewPoint(16, 1),
		math.NewPoint(13, 3),
		math.NewPoint(14, 3),
	}
	assert.Equal(t, expectedOrder, asteroids)
}

func TestVaporizeBigExample(t *testing.T) {
	vaporizer := NewVaporizer(NewMap([]string{
		".#..##.###...#######",
		"##.############..##.",
		".#.######.########.#",
		".###.#######.####.#.",
		"#####.##.#.##.###.##",
		"..#####..#.#########",
		"####################",
		"#.####....###.#.#.##",
		"##.#################",
		"#####.##.###..####..",
		"..######..##.#######",
		"####.##.####...##..#",
		".#####..#.######.###",
		"##...#.##########...",
		"#.##########.#######",
		".####.#.###.###.#.##",
		"....##.##.###..#####",
		".#.#.###########.###",
		"#.#.#.#####.####.###",
		"###.##.####.##.#..##",
	}))
	location := math.NewPoint(11, 13)
	asteroids := vaporizer.Vaporize(location)
	assert.Equal(t, 299, len(asteroids))
	assert.Equal(t, math.NewPoint(11, 12), asteroids[0])
	assert.Equal(t, math.NewPoint(12, 1), asteroids[1])
	assert.Equal(t, math.NewPoint(12, 2), asteroids[2])
	assert.Equal(t, math.NewPoint(12, 8), asteroids[9])
	assert.Equal(t, math.NewPoint(16, 0), asteroids[19])
	assert.Equal(t, math.NewPoint(16, 9), asteroids[49])
	assert.Equal(t, math.NewPoint(10, 16), asteroids[99])
	assert.Equal(t, math.NewPoint(9, 6), asteroids[198])
	assert.Equal(t, math.NewPoint(8, 2), asteroids[199])
	assert.Equal(t, math.NewPoint(10, 9), asteroids[200])
	assert.Equal(t, math.NewPoint(11, 1), asteroids[298])
}
