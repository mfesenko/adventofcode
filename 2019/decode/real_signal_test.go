package decode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRealSignalDecoder(t *testing.T) {
	for _, test := range realSignalTests() {
		message, err := ParseMessage(test.input)
		require.NoError(t, err)

		decoder := NewRealSignalDecoder(test.phaseCount)
		output := decoder.Decode(message)

		assert.Equal(t, test.output, output.String())
	}
}

func realSignalTests() []testData {
	return []testData{
		{
			input:      "03036732577212944063491565474664",
			phaseCount: 100,
			output:     "84462026",
		},
		{
			input:      "02935109699940807407585447034323",
			phaseCount: 100,
			output:     "78725270",
		},
		{
			input:      "03081770884921959731165446850517",
			phaseCount: 100,
			output:     "53553731",
		},
	}
}
