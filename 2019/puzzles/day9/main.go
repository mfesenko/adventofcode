package main

import (
	"fmt"

	"github.com/mfesenko/adventofcode/2019/intcode"
)

const (
	testMode        = int64(1)
	sensorBoostMode = int64(2)
)

func main() {
	program, err := intcode.LoadProgram("input.txt")
	if err != nil {
		fmt.Printf("Failed to load input: %v\n", err)
		return
	}

	fmt.Printf("Part 1: %v\n", runBoostProgram(program, testMode))
	fmt.Printf("Part 2: %v\n", runBoostProgram(program, sensorBoostMode))
}

func runBoostProgram(program *intcode.Program, mode int64) int64 {
	computer := intcode.NewComputer()
	computer.SetProgram(program)
	computer.Input() <- mode
	computer.Execute()

	n := len(computer.Output())
	output := make([]int64, n)
	for i := 0; i < n; i++ {
		output[i] = <-computer.Output()
	}

	if n == 1 {
		return output[0]
	}

	fmt.Printf("Program execution failed: %v\n", output)
	return -1
}
