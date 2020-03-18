package main

import (
	"fmt"

	"github.com/mfesenko/adventofcode/2019/chemistry"
	"github.com/mfesenko/adventofcode/2019/input"
)

func main() {
	input, err := input.LoadFromFile("input.txt")
	if err != nil {
		fmt.Printf("Failed to load input: %v\n", err)
		return
	}

	reactions, err := chemistry.ParseReactions(input)
	if err != nil {
		fmt.Printf("Failed to parse reactions: %v\n", err)
	}

	factory := chemistry.NewNanofactory(reactions, "ORE", "FUEL")
	fmt.Printf("Minimum amount of ORE to produce 1 FUEL: %v\n", factory.ProductCost())
	fmt.Printf("Maximum amount of fuel: %v\n", factory.Produce(1000000000000))
}
