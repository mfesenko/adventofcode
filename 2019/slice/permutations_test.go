package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGivenNilSlicePermutationsInt64ReturnsNil(t *testing.T) {
	assert.Nil(t, PermutationsInt64(nil))
}

func TestGivenEmptySlicePermutationsInt64ReturnsNil(t *testing.T) {
	assert.Nil(t, PermutationsInt64([]int64{}))
}

func TestGivenNonEmptySlicePermutationsInt64ReturnsAllPossiblePermutations(t *testing.T) {
	input := []int64{1, 2, 3}
	expected := [][]int64{
		{1, 2, 3}, {2, 1, 3}, {3, 1, 2},
		{1, 3, 2}, {2, 3, 1}, {3, 2, 1},
	}

	permutations := PermutationsInt64(input)

	require.Equal(t, len(expected), len(permutations))
	for _, p := range expected {
		assert.Contains(t, permutations, p)
	}
}
