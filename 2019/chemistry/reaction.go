package chemistry

import (
	"fmt"
	"strings"
)

type (
	// Reaction describes a chemical reaction
	Reaction struct {
		Input  map[string]Node
		Output Node
	}

	// Node describes a node in the reaction
	Node struct {
		Chemical string
		Amount   int64
	}
)

// NewReaction creates a new reaction
func NewReaction(input map[string]Node, output Node) Reaction {
	return Reaction{
		Input:  input,
		Output: output,
	}
}

// NewNode creates a new node
func NewNode(chemical string, amount int64) Node {
	return Node{
		Chemical: chemical,
		Amount:   amount,
	}
}

// ParseReactions parses reactions from the list of strings
func ParseReactions(input []string) (map[string]Reaction, error) {
	reactions := map[string]Reaction{}
	for _, line := range input {
		reaction, err := parseReaction(line)
		if err != nil {
			return nil, err
		}

		reactions[reaction.Output.Chemical] = reaction
	}
	return reactions, nil
}

func parseReaction(input string) (Reaction, error) {
	parts := strings.Split(input, "=>")
	if len(parts) != 2 {
		return Reaction{}, fmt.Errorf("invalid reaction format")
	}

	inputParts := strings.Split(parts[0], ",")
	inputNodes := map[string]Node{}
	for _, input := range inputParts {
		inputNode, err := parseNode(strings.Trim(input, " "))
		if err != nil {
			return Reaction{}, err
		}
		inputNodes[inputNode.Chemical] = inputNode
	}

	outputNode, err := parseNode(strings.Trim(parts[1], " "))
	if err != nil {
		return Reaction{}, err
	}

	return NewReaction(inputNodes, outputNode), nil
}

func parseNode(input string) (Node, error) {
	var amount int64
	var chemical string
	_, err := fmt.Sscanf(input, "%d %s", &amount, &chemical)
	if err != nil {
		return Node{}, err
	}

	return NewNode(chemical, amount), nil
}
