package decode

const (
	_multiplier = 10000
)

// RealSignalDecoder decodes input message
type RealSignalDecoder struct {
	phaseCount int
}

// NewRealSignalDecoder creates a RealSignalDecoder
func NewRealSignalDecoder(phaseCount int) *RealSignalDecoder {
	return &RealSignalDecoder{
		phaseCount: phaseCount,
	}
}

// Decode decodes input message
func (d *RealSignalDecoder) Decode(input Message) Message {
	n := input.Size()
	realMessage := make([]int8, n*_multiplier)
	for i := 0; i < _multiplier; i++ {
		copy(realMessage[i*n:], input.Data())
	}

	for i := 0; i < d.phaseCount; i++ {
		realMessage = d.applyPhase(realMessage, input.Offset())
	}

	return NewMessage(realMessage[input.Offset():input.Offset()+8], 0)
}

func (d *RealSignalDecoder) applyPhase(input []int8, offset int) []int8 {
	result := make([]int8, len(input))
	sum := int64(0)
	for i := len(input) - 1; i >= offset; i-- {
		sum += int64(input[i])
		result[i] = int8(sum % 10)
	}
	return result
}
