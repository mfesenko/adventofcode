package decode

import (
	"github.com/mfesenko/adventofcode/2019/math"
)

// FFTDecoder decodes input message with the Flawed Frequency Transmission algorithm
type FFTDecoder struct {
	phaseCount int
}

// NewFFTDecoder creates a FFTDecoder
func NewFFTDecoder(phaseCount int) *FFTDecoder {
	return &FFTDecoder{
		phaseCount: phaseCount,
	}
}

// Decode decodes input message with the Flawed Frequency Transmission algorithm
func (d *FFTDecoder) Decode(input Message) Message {
	data := input.Data()
	for i := 0; i < d.phaseCount; i++ {
		data = d.applyPhase(data)
	}
	return NewMessage(data[:8], 0)
}

func (d *FFTDecoder) applyPhase(input []int8) []int8 {
	result := make([]int8, len(input))
	for i := range input {
		result[i] = d.applyPatternForIndex(input, i+1)
	}
	return result
}

func (d *FFTDecoder) applyPatternForIndex(input []int8, i int) int8 {
	result := d.sum(input, i, i) - d.sum(input, 3*i, i)
	return int8(math.Abs(result) % 10)
}

func (d *FFTDecoder) sum(input []int8, start int, batchSize int) int64 {
	result := int64(0)
	max := len(input)
	for i := start; i <= max; i += 4 * batchSize {
		for j := 0; j < batchSize; j++ {
			address := i + j - 1
			if address >= max {
				break
			}
			result += int64(input[address])
		}
	}
	return result
}
