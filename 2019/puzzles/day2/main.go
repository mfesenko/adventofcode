package main

import (
	"fmt"

	"github.com/mfesenko/adventofcode/2019/intcode"
)

func main() {
	program, err := intcode.LoadProgram("input.txt")
	if err != nil {
		fmt.Printf("Failed to load input: %v\n", err)
		return
	}

	computer := intcode.NewComputer()

	fmt.Printf("Part 1: %v\n", runGravityAssistProgram(computer, program, 12, 2))
	fmt.Printf("Part 2: %v\n", findNounAndVerb(computer, program, 19690720))
}

func runGravityAssistProgram(computer *intcode.Computer, program *intcode.Program, noun int64, verb int64) int64 {
	program = program.Copy()
	program.SetNoun(noun)
	program.SetVerb(verb)
	computer.SetProgram(program)
	return computer.Execute()
}

func findNounAndVerb(computer *intcode.Computer, program *intcode.Program, expectedResult int64) int64 {
	for noun := int64(0); noun < 100; noun++ {
		for verb := int64(0); verb < 100; verb++ {
			if runGravityAssistProgram(computer, program, noun, verb) == expectedResult {
				return noun*100 + verb
			}
		}
	}
	return -1
}
