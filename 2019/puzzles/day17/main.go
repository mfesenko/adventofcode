package main

import (
	"fmt"

	"github.com/mfesenko/adventofcode/2019/ascii"
	"github.com/mfesenko/adventofcode/2019/intcode"
)

func main() {
	program, err := intcode.LoadProgram("input.txt")
	if err != nil {
		fmt.Printf("Failed to load input: %v\n", err)
		return
	}

	program.Write(0, 2)
	computer := intcode.NewComputer()
	computer.SetProgram(program)
	robot := ascii.NewRobot(computer)
	alignmentParameters, collectedDust := robot.Run(false)
	fmt.Printf("Alignment parameters: %v\n", alignmentParameters)
	fmt.Printf("Collected dust: %v\n", collectedDust)
}
