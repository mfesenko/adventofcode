package main

import (
	"fmt"
	"sync"

	"github.com/mfesenko/adventofcode/2019/intcode"
	"github.com/mfesenko/adventofcode/2019/slice"
)

type (
	signalGenerator          func() [][]int64
	thrusterSignalCalculator func(program *intcode.Program, phaseSignals []int64) int64
)

func main() {
	program, err := intcode.LoadProgram("input.txt")
	if err != nil {
		fmt.Printf("Failed to load input: %v\n", err)
		return
	}

	thrusterSignal, inputSignals := findMaxThrusterSignal(program, generatePhaseSettings, calculateThrusterSignal)
	fmt.Printf("Part 1 - Thruster signal: %v, input signals: %v \n", thrusterSignal, inputSignals)

	thrusterSignal, inputSignals = findMaxThrusterSignal(program, generatePhaseSettingsWithFeedbackLoop, calculateThrusterSignalWithFeedbackLoop)
	fmt.Printf("Part 2 - Thruster signal: %v, input signals: %v \n", thrusterSignal, inputSignals)
}

func findMaxThrusterSignal(program *intcode.Program, generator signalGenerator, calculator thrusterSignalCalculator) (int64, []int64) {
	maxThrusterSignal := int64(0)
	var maxInputSignal []int64
	signals := generator()
	for _, signal := range signals {
		thrusterSignal := calculator(program, signal)
		if thrusterSignal > maxThrusterSignal {
			maxThrusterSignal = thrusterSignal
			maxInputSignal = signal
		}
	}
	return maxThrusterSignal, maxInputSignal
}

func generatePhaseSettings() [][]int64 {
	return generatePhaseSettingsInRange(0, 4)
}

func generatePhaseSettingsWithFeedbackLoop() [][]int64 {
	return generatePhaseSettingsInRange(5, 9)
}

func generatePhaseSettingsInRange(min int64, max int64) [][]int64 {
	count := max - min + 1
	a := make([]int64, count)
	for i := int64(0); i < count; i++ {
		a[i] = i + min
	}
	return slice.PermutationsInt64(a)
}

func calculateThrusterSignal(program *intcode.Program, phaseSignals []int64) int64 {
	done := doneSignal(len(phaseSignals))
	output := int64(0)
	for _, signal := range phaseSignals {
		computer := intcode.NewComputer()
		computer.SetProgram(program)
		computer.ExecuteAsync(done)
		computer.Input() <- signal
		computer.Input() <- output
		output = <-computer.Output()
	}
	done.Wait()
	return output
}

func calculateThrusterSignalWithFeedbackLoop(program *intcode.Program, phaseSignals []int64) int64 {
	count := len(phaseSignals)
	computers := make([]*intcode.Computer, count)
	for i := 0; i < count; i++ {
		computer := intcode.NewComputer()
		computer.SetProgram(program)
		computers[i] = computer
		if i > 0 {
			computer.SetInput(computers[i-1].Output())
		}
	}
	computers[0].SetInput(computers[count-1].Output())

	done := doneSignal(count)
	for i := count - 1; i >= 0; i-- {
		computer := computers[i]
		computer.Input() <- phaseSignals[i]
		if i == 0 {
			computer.Input() <- 0
		}
		computer.ExecuteAsync(done)
	}
	done.Wait()
	return <-computers[count-1].Output()
}

func doneSignal(count int) *sync.WaitGroup {
	done := &sync.WaitGroup{}
	done.Add(count)
	return done
}
