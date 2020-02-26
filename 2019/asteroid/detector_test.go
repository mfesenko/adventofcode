package asteroid

import (
	"testing"

	"github.com/mfesenko/adventofcode/2019/math"
	"github.com/stretchr/testify/assert"
)

type testData struct {
	input    []string
	count    int
	location math.Point
}

func TestFindBestLocation(t *testing.T) {
	tests := tests()
	for _, test := range tests {
		d := NewDetector(NewMap(test.input))
		location, count := d.FindBestLocation()
		assert.Equal(t, test.location, location)
		assert.Equal(t, test.count, count)
	}
}

func tests() []testData {
	return []testData{
		{
			input: []string{
				".#..#",
				".....",
				"#####",
				"....#",
				"...##",
			},
			count:    8,
			location: math.NewPoint(3, 4),
		},
		{
			input: []string{
				"......#.#.",
				"#..#.#....",
				"..#######.",
				".#.#.###..",
				".#..#.....",
				"..#....#.#",
				"#..#....#.",
				".##.#..###",
				"##...#..#.",
				".#....####",
			},
			count:    33,
			location: math.NewPoint(5, 8),
		},
		{
			input: []string{
				"#.#...#.#.",
				".###....#.",
				".#....#...",
				"##.#.#.#.#",
				"....#.#.#.",
				".##..###.#",
				"..#...##..",
				"..##....##",
				"......#...",
				".####.###.",
			},
			count:    35,
			location: math.NewPoint(1, 2),
		},
		{
			input: []string{
				".#..#..###",
				"####.###.#",
				"....###.#.",
				"..###.##.#",
				"##.##.#.#.",
				"....###..#",
				"..#.#..#.#",
				"#..#.#.###",
				".##...##.#",
				".....#.#..",
			},
			count:    41,
			location: math.NewPoint(6, 3),
		},
		{
			input: []string{
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
			},
			count:    210,
			location: math.NewPoint(11, 13),
		}}
}
