package chemistry

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGivenInvalidInputParseReactionsReturnsAnError(t *testing.T) {
	input := []string{
		"10 ORE => 10 A",
		"asdf",
	}
	reactions, err := ParseReactions(input)

	assert.Error(t, err)
	assert.Nil(t, reactions)
}

func TestGivenValidInputParseReactionsReturnsReactions(t *testing.T) {
	input := []string{
		"10 ORE => 10 A",
		"2 ORE, 1 A => 1 B",
	}
	reactions, err := ParseReactions(input)

	assert.NoError(t, err)
	require.Equal(t, 2, len(reactions))

	reactionA := NewReaction(
		map[string]Node{
			"ORE": NewNode("ORE", 10),
		},
		NewNode("A", 10),
	)
	assert.Equal(t, reactionA, reactions["A"])

	reactionB := NewReaction(
		map[string]Node{
			"ORE": NewNode("ORE", 2),
			"A":   NewNode("A", 1),
		},
		NewNode("B", 1),
	)
	assert.Equal(t, reactionB, reactions["B"])
}

func TestGivenInvalidInputStringParseReactionReturnsAnError(t *testing.T) {
	tests := []string{
		"asdf",
		"asdf => 1 qwerty",
		"1 asdf => qwerty",
	}
	for _, test := range tests {
		_, err := parseReaction(test)
		assert.Error(t, err)
	}
}

func TestGivenValidInputStringParseReactionReturnsAReaction(t *testing.T) {
	reaction, err := parseReaction("2 AB, 3 BC, 4 CA => 1 FUEL")

	assert.NoError(t, err)
	assert.Equal(t, NewNode("FUEL", 1), reaction.Output)
	require.Equal(t, 3, len(reaction.Input))
	assert.Equal(t, NewNode("AB", 2), reaction.Input["AB"])
	assert.Equal(t, NewNode("BC", 3), reaction.Input["BC"])
	assert.Equal(t, NewNode("CA", 4), reaction.Input["CA"])
}
