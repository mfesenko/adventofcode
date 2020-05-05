package ascii

import (
	"strings"
)

const (
	maxFunctionLen        = 20
	minCommandsInFunction = 4
	commandSeparator      = ","
)

// MainRoutineBuilder breaks down the series of commands into functions that will be called by a main routine
type MainRoutineBuilder struct {
	routine   string
	functions map[string]string
}

// NewMainRoutineBuilder creates a MainRoutineBuilder
func NewMainRoutineBuilder() *MainRoutineBuilder {
	return &MainRoutineBuilder{}
}

// MainRoutine returns a main routine
func (b *MainRoutineBuilder) MainRoutine() string {
	return b.routine
}

// Function returns a function definition
func (b *MainRoutineBuilder) Function(functionName string) string {
	return b.functions[functionName]
}

// Build breaks down the series of commands into functions that will be called by a main routine
func (b *MainRoutineBuilder) Build(commands []string, functionNames []string) {
	b.functions = map[string]string{}
	b.routine = strings.Join(commands, commandSeparator)

	for _, name := range functionNames {
		function, offset := b.baseFunction(commands)
		minCount := b.count(function)

		for i := offset; i < len(commands)-1; i = i + 2 {
			firstCommand := commands[i]
			secondCommand := commands[i+1]
			if b.isFunction(firstCommand) {
				break
			}

			if b.isFunction(secondCommand) {
				break
			}

			candidate := function + commandSeparator + firstCommand + commandSeparator + secondCommand
			if len(candidate) > maxFunctionLen {
				break
			}

			candidateCount := b.count(candidate)
			if candidateCount < minCount {
				break
			}

			function = candidate
		}

		b.functions[name] = function
		b.routine = strings.ReplaceAll(b.routine, function, name)
		commands = strings.Split(b.routine, commandSeparator)
	}
}

func (b *MainRoutineBuilder) isFunction(command string) bool {
	_, ok := b.functions[command]
	return ok
}

func (b *MainRoutineBuilder) count(function string) int {
	return strings.Count(b.routine, function)
}

func (b *MainRoutineBuilder) baseFunction(commands []string) (string, int) {
	functionBuilder := &strings.Builder{}
	count := 0
	start := 0
	for i, command := range commands {
		if b.isFunction(command) {
			continue
		}

		if count > 0 {
			functionBuilder.WriteString(commandSeparator)
		}

		functionBuilder.WriteString(command)
		count++
		if count == minCommandsInFunction {
			start = i + 1
			break
		}
	}
	return functionBuilder.String(), start
}
