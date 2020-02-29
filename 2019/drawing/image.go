package drawing

import (
	"strings"
)

const (
	// Black represents black pixel
	Black = uint16(0)
	// White represents white pixel
	White = uint16(1)
	// Transparent represents transparent pixel
	Transparent = uint16(2)
)

// Image respresents an image
type Image struct {
	width  int
	height int
	layers [][][]uint16
}

// NewImage creates a new image
func NewImage(width int, height int, layerCount int) *Image {
	layers := make([][][]uint16, layerCount)
	for i := 0; i < layerCount; i++ {
		layer := make([][]uint16, height)
		for j := 0; j < height; j++ {
			layer[j] = make([]uint16, width)
		}
		layers[i] = layer
	}

	return &Image{
		width:  width,
		height: height,
		layers: layers,
	}
}

// CheckSum returns a checksum for the image that is calculated as the number of 1 digits multiplied
// by the number of 2 digits in the layer of the image that contains the fewest 0 digits.
func (i *Image) CheckSum() int {
	zeroPerLayer := i.valueCountPerLayer(Black)
	minCount := zeroPerLayer[0]
	minLayer := 0
	for i := 1; i < len(zeroPerLayer); i++ {
		if minCount > zeroPerLayer[i] {
			minCount = zeroPerLayer[i]
			minLayer = i
		}
	}

	return i.valueCountInLayer(White, minLayer) * i.valueCountInLayer(Transparent, minLayer)
}

func (i *Image) valueCountPerLayer(value uint16) []int {
	counts := make([]int, len(i.layers))
	for j := range i.layers {
		counts[j] = i.valueCountInLayer(value, j)
	}
	return counts
}

func (i *Image) valueCountInLayer(value uint16, layer int) int {
	count := 0
	for _, row := range i.layers[layer] {
		for _, x := range row {
			if x == value {
				count++
			}
		}
	}
	return count
}

// GetPixel returns the value of the pixel for given coordinates
func (i *Image) GetPixel(x int, y int, layer int) uint16 {
	return i.layers[layer][y][x]
}

// SetPixel sets the value of the pixel for given coordinates
func (i *Image) SetPixel(x int, y int, layer int, value uint16) {
	i.layers[layer][y][x] = value
}

// Render renders the image
func (i *Image) Render() [][]uint16 {
	rendered := make([][]uint16, i.height)
	for j := 0; j < i.height; j++ {
		row := make([]uint16, i.width)
		for k := 0; k < i.width; k++ {
			row[k] = i.getPixel(j, k)
		}
		rendered[j] = row
	}
	return rendered
}

func (i *Image) getPixel(j int, k int) uint16 {
	for _, layer := range i.layers {
		pixel := layer[j][k]
		if pixel != Transparent {
			return pixel
		}
	}
	return Transparent
}

// String converts rendered image to a string
func (i *Image) String() string {
	builder := &strings.Builder{}
	rendered := i.Render()
	for _, row := range rendered {
		for _, pixel := range row {
			pixelRune := ' '
			if pixel == White {
				pixelRune = '★'
			}
			builder.WriteRune((pixelRune))
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}

// ParseImage parses image from the string representation
func ParseImage(data string, width int, height int) *Image {
	length := len(data)
	layerCount := length / (width * height)
	image := NewImage(width, height, layerCount)
	for i := 0; i < length; i++ {
		x := i % width
		y := (i % (height * width)) / width
		l := i / (height * width)
		value := uint16(data[i] - '0')
		image.SetPixel(x, y, l, value)
	}
	return image
}
