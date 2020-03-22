package main

import (
	"fmt"

	"github.com/mfesenko/adventofcode/2019/intcode"
	"github.com/mfesenko/adventofcode/2019/repair"
)

func main() {
	program, err := intcode.LoadProgram("input.txt")
	if err != nil {
		fmt.Printf("Failed to load input: %v\n", err)
		return
	}

	computer := intcode.NewComputer()
	computer.SetProgram(program)
	droid := repair.NewDroid(computer)
	droid.Explore()
	fmt.Printf("Shortest path to oxygen system: %v\n", droid.ShortestPathToOxygenSystem())
	fmt.Printf("Minutes to fill area with oxygen: %v\n", droid.MinutesToFillWithOxygen())
}
