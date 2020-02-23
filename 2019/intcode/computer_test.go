package intcode

import (
	"fmt"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestExitCodeReturnsFirstCodeFromTheProgram(t *testing.T) {
	program := randomProgram()
	computerWithProgram(program, func(computer *Computer) {
		assert.Equal(t, program.Read(0), computer.ExitCode())
	})
}

func TestExecuteUpdatesTheProgram(t *testing.T) {
	type testData struct {
		name          string
		startProgram  *Program
		resultProgram *Program
	}
	tests := []testData{
		{
			name:          "halt in the middle of the program",
			startProgram:  NewProgram([]int64{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}),
			resultProgram: NewProgram([]int64{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}),
		},
		{
			name:          "3 * 2 = 6",
			startProgram:  NewProgram([]int64{2, 3, 0, 3, 99}),
			resultProgram: NewProgram([]int64{2, 3, 0, 6, 99}),
		},
		{
			name:          "1 + 1 = 2",
			startProgram:  NewProgram([]int64{1, 0, 0, 0, 99}),
			resultProgram: NewProgram([]int64{2, 0, 0, 0, 99}),
		},
		{
			name:          "99 * 99 = 9801",
			startProgram:  NewProgram([]int64{2, 4, 4, 5, 99, 0}),
			resultProgram: NewProgram([]int64{2, 4, 4, 5, 99, 9801}),
		},
		{
			name:          "halt in the end of the program",
			startProgram:  NewProgram([]int64{1, 1, 1, 4, 99, 5, 6, 0, 99}),
			resultProgram: NewProgram([]int64{30, 1, 1, 4, 2, 5, 6, 0, 99}),
		},
		{
			name:          "multiply with position and immediate mode",
			startProgram:  NewProgram([]int64{1002, 4, 3, 4, 33}),
			resultProgram: NewProgram([]int64{1002, 4, 3, 4, 99}),
		},
		{
			name:          "add with immediate mode with negative numbers",
			startProgram:  NewProgram([]int64{1101, 100, -1, 4, 0}),
			resultProgram: NewProgram([]int64{1101, 100, -1, 4, 99}),
		},
		{
			name:          "add with relative mode",
			startProgram:  NewProgram([]int64{109, -3, 22101, 15, 4, 5, 99}),
			resultProgram: NewProgram([]int64{109, -3, 12, 15, 4, 5, 99}),
		},
	}
	for _, test := range tests {
		computerWithProgram(test.startProgram, func(computer *Computer) {
			computer.Execute()

			assert.Equal(t, test.resultProgram, computer.program, "test name: \"%v\"", test.name)
		})
	}
}

func TestExecuteWithCustomInputAndOutput(t *testing.T) {
	computerWithProgram(NewProgram([]int64{3, 0, 4, 0, 4, 0, 99}), func(computer *Computer) {
		input := make(chan int64, 1)
		output := make(chan int64, 2)
		computer.SetInput(input)
		computer.SetOutput(output)

		inputValue := randomValueBelowHundred()
		input <- inputValue
		computer.Execute()

		assert.Equal(t, 2, len(output))
		assert.Equal(t, inputValue, <-output)
		assert.Equal(t, inputValue, <-output)
	})
}

func TestExecuteWithInputProducesExpectedOutput(t *testing.T) {
	type testData struct {
		name        string
		program     *Program
		inputOutput map[int64]int64
	}
	tests := []testData{
		{
			name:    "output 1 if input equals to 8 with position mode",
			program: NewProgram([]int64{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}),
			inputOutput: map[int64]int64{
				3:  0,
				8:  1,
				23: 0,
			},
		},
		{
			name:    "output 1 if input equals to 8 with immediate mode",
			program: NewProgram([]int64{3, 3, 1108, -1, 8, 3, 4, 3, 99}),
			inputOutput: map[int64]int64{
				3:  0,
				8:  1,
				23: 0,
			},
		},
		{
			name:    "output 1 if input is less then 8 with position mode",
			program: NewProgram([]int64{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}),
			inputOutput: map[int64]int64{
				3:  1,
				8:  0,
				23: 0,
			},
		},
		{
			name:    "output 1 if input is less then 8 with immediate mode",
			program: NewProgram([]int64{3, 3, 1107, -1, 8, 3, 4, 3, 99}),
			inputOutput: map[int64]int64{
				3:  1,
				8:  0,
				23: 0,
			},
		},
		{
			name:    "output 0 if input is 0 with position mode",
			program: NewProgram([]int64{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}),
			inputOutput: map[int64]int64{
				0:  0,
				17: 1,
			},
		},
		{
			name:    "output 0 if input is 0 with immediate mode",
			program: NewProgram([]int64{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}),
			inputOutput: map[int64]int64{
				0:  0,
				17: 1,
			},
		},
		{
			name: "output 999 if less then 8, 1000 if equal to 8, 1001 if greater than 8",
			program: NewProgram([]int64{
				3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
			}),
			inputOutput: map[int64]int64{
				2:  999,
				8:  1000,
				10: 1001,
			},
		},
	}
	for _, test := range tests {
		for input, output := range test.inputOutput {
			computerWithProgram(test.program, func(computer *Computer) {
				computer.Input() <- input

				computer.Execute()

				assert.Equal(t, 1, len(computer.Output()), "test name: \"%v\", input: %v", test.name, input)
				assert.Equal(t, output, <-computer.Output(), "test name: \"%v\", input: %v", test.name, input)
			})
		}
	}
}

func TestAdjustRelativeBaseChangesRelativeBase(t *testing.T) {
	type testData struct {
		name         string
		program      *Program
		relativeBase int64
	}
	tests := []testData{
		{
			name:         "adjust relative base with immediate mode",
			program:      NewProgram([]int64{109, 19, 109, -2, 99}),
			relativeBase: 17,
		},
		{
			name:         "adjust relative base with positional mode",
			program:      NewProgram([]int64{9, 2, 99}),
			relativeBase: 99,
		},
	}
	for _, test := range tests {
		computerWithProgram(test.program, func(computer *Computer) {
			computer.Execute()

			assert.Equal(t, test.relativeBase, computer.relativeBase, "test: \"%v\"", test.name)
		})
	}
}

func TestCopyProgramToOutput(t *testing.T) {
	program := NewProgram([]int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99})
	computerWithProgram(program, func(computer *Computer) {
		size := program.Len()
		output := make(chan int64, size)
		computer.SetOutput(output)

		computer.Execute()

		assert.Equal(t, size, int64(len(output)))
		result := make([]int64, size)
		for i := int64(0); i < size; i++ {
			result[i] = <-output
		}
		assert.Equal(t, program, NewProgram(result))
	})
}

func TestOutput16DigitNumber(t *testing.T) {
	program := NewProgram([]int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0})
	computerWithProgram(program, func(computer *Computer) {
		computer.Execute()

		result := <-computer.Output()
		resultStr := strconv.FormatInt(result, 10)
		assert.Equal(t, 16, len(resultStr))
	})
}

func TestOutputLargeNumber(t *testing.T) {
	program := NewProgram([]int64{104, 1125899906842624, 99})
	computerWithProgram(program, func(computer *Computer) {
		computer.Execute()

		assert.Equal(t, program.Read(1), <-computer.Output())
	})
}

func TestExecuteAsync(t *testing.T) {
	program := NewProgram([]int64{104, 1125899906842624, 99})
	computerWithProgram(program, func(computer *Computer) {
		defer goleak.VerifyNone(t)
		done := &sync.WaitGroup{}
		done.Add(1)
		computer.ExecuteAsync(done)
		done.Wait()

		assert.Equal(t, program.Read(1), <-computer.Output())
	})
}

func TestExecutePanicsWhenTryingToReadParameterWithUnsupportedParameterMode(t *testing.T) {
	program := NewProgram([]int64{501, 1, 1, 1})
	computerWithProgram(program, func(computer *Computer) {
		assert.PanicsWithValue(t, fmt.Sprintf(_unsupportedParameterModeForReading, 5), func() {
			computer.Execute()
		})
	})
}

func TestExecutePanicsWhenTryingToWriteParameterWithUnsupportedParameterMode(t *testing.T) {
	program := NewProgram([]int64{11101, 1, 1, 1})
	computerWithProgram(program, func(computer *Computer) {
		assert.PanicsWithValue(t, fmt.Sprintf(_unsupportedParameterModeForWriting, 1), func() {
			computer.Execute()
		})
	})
}

func computerWithProgram(program *Program, test func(computer *Computer)) {
	computer := NewComputer()
	computer.SetProgram(program)
	test(computer)
}
