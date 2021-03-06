package drawing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetPixel(t *testing.T) {
	image := NewImage(4, 4, 1)
	assert.Equal(t, Black, image.GetPixel(2, 3, 0))

	image.SetPixel(2, 3, 0, White)

	assert.Equal(t, White, image.GetPixel(2, 3, 0))
}

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
	height := 3
	image := ParseImage("022222112222221222000022", width, height)

	assert.Equal(t, [][]uint16{
		{Black, White},
		{White, Black},
		{Transparent, Transparent},
	}, image.Render())
}

func TestString(t *testing.T) {
	width := 2
	height := 2
	image := ParseImage("0222112222120000", width, height)

	assert.Equal(t, " ★\n★ \n", image.String())
}
