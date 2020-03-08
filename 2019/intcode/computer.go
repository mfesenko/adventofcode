package intcode

import (
	"fmt"
)

const (
	_unsupportedParameterModeForReading = "unsupported parameter mode for reading: %v"
	_unsupportedParameterModeForWriting = "unsupported parameter mode for writing: %v"
)

// Computer executes a program
type Computer struct {
	program      *Program
	input        chan int64
	output       chan int64
	relativeBase int64
}

// NewComputer creates a computer
func NewComputer() *Computer {
	return &Computer{
		input:        make(chan int64, 1),
		output:       make(chan int64, 10),
		relativeBase: 0,
	}
}

// SetProgram sets program for the computer
func (c *Computer) SetProgram(program *Program) {
	c.program = program.Copy()
}

// SetInput sets input channel for the computer
func (c *Computer) SetInput(input chan int64) {
	c.input = input
}

// Input returns the input channel for the computer
func (c *Computer) Input() chan int64 {
	return c.input
}

// SetOutput sets output channel for the computer
func (c *Computer) SetOutput(output chan int64) {
	c.output = output
}

// Output returns the output channel for the computer
func (c *Computer) Output() chan int64 {
	return c.output
}

// ExitCode returns an exit code from the program
func (c *Computer) ExitCode() int64 {
	return c.program.Read(0)
}

// Execute executes the Intcode program
func (c *Computer) Execute() {
	for i := int64(0); i < c.program.Len(); {
		i = c.processCommand(i)
	}
	close(c.output)
}

func (c *Computer) processCommand(i int64) int64 {
	operation := ParseOperation(i, c.program.Read(i))
	switch operation.operationType {
	case AddOperation:
		c.binaryOperation(operation, func(a int64, b int64) int64 {
			return a + b
		})

	case MultiplyOperation:
		c.binaryOperation(operation, func(a int64, b int64) int64 {
			return a * b
		})

	case LessThanOperation:
		c.binaryOperation(operation, func(a int64, b int64) int64 {
			if a < b {
				return 1
			}
			return 0
		})

	case EqualsOperation:
		c.binaryOperation(operation, func(a int64, b int64) int64 {
			if a == b {
				return 1
			}
			return 0
		})

	case ReadInputOperation:
		c.writeParameter(operation, 0, <-c.input)

	case WriteOutputOperation:
		c.output <- c.readParameter(operation, 0)

	case AdjustRelativeBaseOperation:
		c.relativeBase = c.relativeBase + c.readParameter(operation, 0)

	case JumpIfTrueOperation:
		if c.readParameter(operation, 0) != 0 {
			return c.readParameter(operation, 1)
		}

	case JumpIfFalseOperation:
		if c.readParameter(operation, 0) == 0 {
			return c.readParameter(operation, 1)
		}

	case HaltOperation:
		return c.program.Len()
	}

	return i + operation.operationType.ParameterCount() + 1
}

func (c *Computer) binaryOperation(operation Operation, op func(a, b int64) int64) {
	a := c.readParameter(operation, 0)
	b := c.readParameter(operation, 1)
	c.writeParameter(operation, 2, op(a, b))
}

func (c *Computer) readParameter(operation Operation, parameterIndex int64) int64 {
	address := operation.address + parameterIndex + 1
	mode := operation.parameterModes[parameterIndex]
	if mode == PositionMode {
		return c.program.Read(c.program.Read(address))
	}

	if mode == ImmediateMode {
		return c.program.Read(address)
	}

	if mode == RelativeMode {
		return c.program.Read(c.relativeBase + c.program.Read(address))
	}

	panic(fmt.Sprintf(_unsupportedParameterModeForReading, mode))
}

func (c *Computer) writeParameter(operation Operation, parameterIndex int64, value int64) {
	address := c.program.Read(operation.address + parameterIndex + 1)
	mode := operation.parameterModes[parameterIndex]
	if mode == PositionMode {
		c.program.Write(address, value)
		return
	}

	if mode == RelativeMode {
		c.program.Write(c.relativeBase+address, value)
		return
	}

	panic(fmt.Sprintf(_unsupportedParameterModeForWriting, mode))
}
