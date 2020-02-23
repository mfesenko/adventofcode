package image

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckSum(t *testing.T) {
	width := 2
	height := 2
	image := ParseImage("0222112222120000", width, height)

	assert.Equal(t, 4, image.CheckSum())
}

func TestParseImage(t *testing.T) {
	width := 3
	height := 2
	image := ParseImage("123456789012", width, height)

	assert.Equal(t, width, image.width)
	assert.Equal(t, height, image.height)
	assert.Equal(t, 2, len(image.layers))
	assert.Equal(t, [][]uint16{
		{1, 2, 3},
		{4, 5, 6},
	}, image.layers[0])
	assert.Equal(t, [][]uint16{
		{7, 8, 9},
		{0, 1, 2},
	}, image.layers[1])
}

func TestRender(t *testing.T) {
	width := 2
	height := 2
	image := ParseImage("0222112222120000", width, height)

	assert.Equal(t, [][]uint16{
		{_black, _white},
		{_white, _black},
	}, image.Render())
}

func TestString(t *testing.T) {
	width := 2
	height := 2
	image := ParseImage("0222112222120000", width, height)

	assert.Equal(t, " ★\n★ \n", image.String())
}
