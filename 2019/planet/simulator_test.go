package planet

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimulate(t *testing.T) {
	tests := tests()
	for _, test := range tests {
		moons, err := ParseMoons(test.input)
		require.NoError(t, err)

		simulator := NewSimulator(moons)
		simulator.Simulate(test.steps)
		assert.Equal(t, test.totalEnergy, simulator.TotalEnergy())
	}
}
