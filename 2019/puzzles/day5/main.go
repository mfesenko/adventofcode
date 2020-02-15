package main

import (
	"fmt"

	"github.com/mfesenko/adventofcode/2019/intcode"
)

const (
	_airConditionerUnitID        = int64(1)
	_thermalRadiatorControllerID = int64(5)
)

func main() {
	program, err := intcode.LoadProgram("input.txt")
	if err != nil {
		fmt.Printf("Failed to load input: %v\n", err)
		return
	}

	fmt.Printf("Part 1: %v\n", runDiagnosticProgram(program, _airConditionerUnitID))
	fmt.Printf("Part 2: %v\n", runDiagnosticProgram(program, _thermalRadiatorControllerID))
}

func runDiagnosticProgram(program *intcode.Program, systemID int64) int64 {
	computer := intcode.NewComputer()
	computer.SetProgram(program)
	computer.Input() <- systemID
	computer.Execute()

	n := len(computer.Output())
	var lastCode int64
	for i := 0; i < n; i++ {
		if lastCode != 0 {
			fmt.Println("Diagnostic tests failed!")
		}
		lastCode = <-computer.Output()
	}
	return lastCode
}
