package image

import (
	"strings"
)

const (
	_black       = uint16(0)
	_white       = uint16(1)
	_transparent = uint16(2)
)

// Image respresents an image
type Image struct {
	width  int
	height int
	layers [][][]uint16
}

// CheckSum returns a checksum for the image that is calculated as the number of 1 digits multiplied
// by the number of 2 digits in the layer of the image that contains the fewest 0 digits.
func (i *Image) CheckSum() int {
	zeroPerLayer := i.digitCountPerLayer(_black)
	minCount := zeroPerLayer[0]
	minLayer := 0
	for i := 1; i < len(zeroPerLayer); i++ {
		if minCount > zeroPerLayer[i] {
			minCount = zeroPerLayer[i]
			minLayer = i
		}
	}

	return i.digitCountInLayer(_white, minLayer) * i.digitCountInLayer(_transparent, minLayer)
}

func (i *Image) digitCountPerLayer(digit uint16) []int {
	counts := make([]int, len(i.layers))
	for j := range i.layers {
		counts[j] = i.digitCountInLayer(digit, j)
	}
	return counts
}

func (i *Image) digitCountInLayer(digit uint16, layer int) int {
	count := 0
	for _, row := range i.layers[layer] {
		for _, x := range row {
			if x == digit {
				count++
			}
		}
	}
	return count
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
		if pixel != _transparent {
			return pixel
		}
	}
	return _transparent
}

// String converts rendered image to a string
func (i *Image) String() string {
	builder := &strings.Builder{}
	rendered := i.Render()
	for _, row := range rendered {
		for _, pixel := range row {
			pixelRune := ' '
			if pixel == _white {
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
	var layers [][][]uint16
	for i := 0; i < len(data); {
		layer := make([][]uint16, height)
		for j := 0; j < height; j++ {
			row := make([]uint16, width)
			for k := 0; k < width; k++ {
				row[k] = uint16(data[i] - '0')
				i++
			}
			layer[j] = row
		}
		layers = append(layers, layer)
	}

	return &Image{
		width:  width,
		height: height,
		layers: layers,
	}
}
