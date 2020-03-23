package decode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testData struct {
	input      string
	phaseCount int
	output     string
}

func TestFFTDecoder(t *testing.T) {
	for _, test := range fftTests() {
		message, err := ParseMessage(test.input)
		require.NoError(t, err)

		decoder := NewFFTDecoder(test.phaseCount)
		output := decoder.Decode(message)

		assert.Equal(t, test.output, output.String())
	}
}

func fftTests() []testData {
	return []testData{
		{
			input:      "12345678",
			phaseCount: 1,
			output:     "48226158",
		},
		{
			input:      "12345678",
			phaseCount: 2,
			output:     "34040438",
		},
		{
			input:      "12345678",
			phaseCount: 3,
			output:     "03415518",
		},
		{
			input:      "12345678",
			phaseCount: 4,
			output:     "01029498",
		},
		{
			input:      "80871224585914546619083218645595",
			phaseCount: 100,
			output:     "24176176",
		},
		{
			input:      "19617804207202209144916044189917",
			phaseCount: 100,
			output:     "73745418",
		},
		{
			input:      "69317163492948606335995924319873",
			phaseCount: 100,
			output:     "52432133",
		},
	}
}
