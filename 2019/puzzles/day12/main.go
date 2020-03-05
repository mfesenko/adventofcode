package main

import (
	"fmt"

	"github.com/mfesenko/adventofcode/2019/input"
	"github.com/mfesenko/adventofcode/2019/planet"
)

func main() {
	input, err := input.LoadFromFile("input.txt")
	if err != nil {
		fmt.Printf("Failed to load input: %v\n", err)
		return
	}

	moons, err := planet.ParseMoons(input)
	if err != nil {
		fmt.Printf("Failed to parse moons: %v\n", err)
		return
	}

	simulator := planet.NewSimulator(moons)
	simulator.Simulate(1000)
	fmt.Printf("Total energy of the system after 1000 steps: %v\n", simulator.TotalEnergy())

	repeater := planet.NewStateRepeater(moons)
	fmt.Printf("Amount of steps until the state is repeated: %v\n", repeater.FindStepCountForRepeat())
}
