package main

import (
	"testing"

	"github.com/mfesenko/adventofcode/2019/intcode"
	"github.com/stretchr/testify/assert"
)

type testData struct {
	program        *intcode.Program
	signals        []int64
	thrusterSignal int64
}

func TestCalculateThrusterSignal(t *testing.T) {
	tests := testsWithoutFeedbackLoop()
	for _, test := range tests {
		signal := calculateThrusterSignal(test.program, test.signals)
		assert.Equal(t, test.thrusterSignal, signal)
	}
}

func TestFindMaxThrusterSignal(t *testing.T) {
	tests := testsWithoutFeedbackLoop()
	for _, test := range tests {
		thrusterSignal, inputSignals := findMaxThrusterSignal(test.program, generatePhaseSettings, calculateThrusterSignal)
		assert.Equal(t, test.thrusterSignal, thrusterSignal)
		assert.Equal(t, test.signals, inputSignals)
	}
}

func testsWithoutFeedbackLoop() []testData {
	return []testData{
		{
			program:        intcode.NewProgram([]int64{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}),
			signals:        []int64{4, 3, 2, 1, 0},
			thrusterSignal: 43210,
		},
		{
			program: intcode.NewProgram([]int64{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23,
				101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}),
			signals:        []int64{0, 1, 2, 3, 4},
			thrusterSignal: 54321,
		},
		{
			program: intcode.NewProgram([]int64{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
				1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}),
			signals:        []int64{1, 0, 4, 3, 2},
			thrusterSignal: 65210,
		},
	}
}
func TestCalculateThrusterSignalWithFeedbackLoop(t *testing.T) {
	tests := testsWithFeedbackLoop()
	for _, test := range tests {
		signal := calculateThrusterSignalWithFeedbackLoop(test.program, test.signals)
		assert.Equal(t, test.thrusterSignal, signal)
	}
}

func TestFindMaxThrusterSignalWithFeedbackLoop(t *testing.T) {
	tests := testsWithFeedbackLoop()
	for _, test := range tests {
		thrusterSignal, inputSignals := findMaxThrusterSignal(test.program, generatePhaseSettingsWithFeedbackLoop, calculateThrusterSignalWithFeedbackLoop)
		assert.Equal(t, test.thrusterSignal, thrusterSignal)
		assert.Equal(t, test.signals, inputSignals)
	}
}

func testsWithFeedbackLoop() []testData {
	return []testData{
		{
			program: intcode.NewProgram([]int64{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
				27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}),
			signals:        []int64{9, 8, 7, 6, 5},
			thrusterSignal: 139629729,
		},
		{
			program: intcode.NewProgram([]int64{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54,
				-5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4,
				53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}),
			signals:        []int64{9, 7, 8, 5, 6},
			thrusterSignal: 18216,
		},
	}
}
