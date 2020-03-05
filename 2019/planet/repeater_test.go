package planet

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStepRepeater(t *testing.T) {
	tests := tests()
	for _, test := range tests {
		moons, err := ParseMoons(test.input)
		require.NoError(t, err)

		repeater := NewStateRepeater(moons)
		steps := repeater.FindStepCountForRepeat()
		assert.Equal(t, test.repeatSteps, steps)
	}
}
