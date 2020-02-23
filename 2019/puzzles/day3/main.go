package main

import (
	"fmt"

	"github.com/mfesenko/adventofcode/2019/input"
)

func main() {
	input, err := input.LoadFromFile("input.txt")
	if err != nil {
		fmt.Printf("Failed to load input: %v\n", err)
		return
	}

	wires, err := parseWires(input)
	if err != nil {
		fmt.Printf("Failed to parse wires: %v\n", err)
		return
	}

	fmt.Printf("Part 1: %v\n", partOne(wires))
	fmt.Printf("Part 2: %v\n", partTwo(wires))
}

func partOne(wires []wire) int32 {
	return findBestIntersection(wires, compareManhattanDistance).manhattanDistance
}

func partTwo(wires []wire) int32 {
	return findBestIntersection(wires, compareStepCount).stepCount
}
