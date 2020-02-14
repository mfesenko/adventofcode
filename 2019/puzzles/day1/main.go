package main

import (
	"fmt"
	"strconv"

	"github.com/mfesenko/adventofcode/2019/input"
)

func main() {
	masses, err := readModuleMasses("input.txt")
	if err != nil {
		fmt.Printf("Failed to load input: %v\n", err)
		return
	}

	fmt.Printf("Part 1: %v\n", calculateFuelMass(masses, naiveFuelCalculator))
	fmt.Printf("Part 2: %v\n", calculateFuelMass(masses, smartFuelCalculator))
}

func readModuleMasses(filePath string) ([]int64, error) {
	data, err := input.LoadFromFile(filePath)
	if err != nil {
		return nil, err
	}

	input := make([]int64, len(data))
	for i, line := range data {
		value, err := strconv.ParseInt(string(line), 10, 64)
		if err != nil {
			return nil, err
		}
		input[i] = value
	}

	return input, nil
}

func calculateFuelMass(masses []int64, calculator func(int64) int64) int64 {
	var fuel int64
	for _, mass := range masses {
		fuel += calculator(mass)
	}
	return fuel
}

func naiveFuelCalculator(mass int64) int64 {
	return mass/3 - 2
}

func smartFuelCalculator(mass int64) int64 {
	fuelTotal := int64(0)
	masses := []int64{mass}
	for i := 0; i < len(masses); i++ {
		fuel := naiveFuelCalculator(masses[i])
		if fuel > 0 {
			fuelTotal += fuel
			masses = append(masses, fuel)
		}
	}
	return fuelTotal
}
