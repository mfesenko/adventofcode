package ascii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsASCII(t *testing.T) {
	tests := map[int64]bool{
		-2:  false,
		0:   true,
		3:   true,
		127: true,
		128: false,
		367: false,
	}
	for input, expected := range tests {
		assert.Equal(t, expected, isASCII(input))
	}
}
