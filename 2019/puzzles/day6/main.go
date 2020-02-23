package main

import (
	"fmt"

	"github.com/mfesenko/adventofcode/2019/input"
	"github.com/mfesenko/adventofcode/2019/navigation"
)

func main() {
	input, err := input.LoadFromFile("input.txt")
	if err != nil {
		fmt.Printf("Failed to load input: %v\n", err)
		return
	}

	orbitMap := navigation.LoadOrbitMap(input)
	fmt.Printf("Part 1: Orbit Map checksum = %v\n", orbitMap.CheckSum())
	fmt.Printf("Part 2: Shortest path between me and Santa = %v\n", orbitMap.FindShortestPath(orbitMap.FindParent("YOU"), orbitMap.FindParent("SAN")))
}
