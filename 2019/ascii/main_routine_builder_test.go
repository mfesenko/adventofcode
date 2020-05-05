package ascii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mainRoutineBuilderTest struct {
	commands      []string
	functionNames []string
	mainRoutine   string
	functions     map[string]string
}

func TestMainRoutineBuilder(t *testing.T) {
	for _, test := range mainRoutineBuilderTests() {
		builder := NewMainRoutineBuilder()
		builder.Build(test.commands, test.functionNames)

		assert.Equal(t, test.mainRoutine, builder.MainRoutine())
		for name, function := range test.functions {
			assert.Equal(t, function, builder.Function(name))
		}
	}
}

func mainRoutineBuilderTests() []mainRoutineBuilderTest {
	return []mainRoutineBuilderTest{
		{
			commands:      []string{"R", "8", "R", "8", "R", "4", "R", "4", "R", "8", "L", "6", "L", "2", "R", "4", "R", "4", "R", "8", "R", "8", "R", "8", "L", "6", "L", "2"},
			functionNames: []string{"X", "Y", "Z"},
			mainRoutine:   "X,Y,Z,Y,X,Z",
			functions: map[string]string{
				"X": "R,8,R,8",
				"Y": "R,4,R,4",
				"Z": "R,8,L,6,L,2",
			},
		},
		{
			commands:      []string{"L", "6", "R", "8", "R", "12", "L", "6", "L", "8", "L", "10", "L", "8", "R", "12", "L", "6", "R", "8", "R", "12", "L", "6", "L", "8", "L", "8", "L", "10", "L", "6", "L", "6", "L", "10", "L", "8", "R", "12", "L", "8", "L", "10", "L", "6", "L", "6", "L", "10", "L", "8", "R", "12", "L", "6", "R", "8", "R", "12", "L", "6", "L", "8", "L", "8", "L", "10", "L", "6", "L", "6", "L", "10", "L", "8", "R", "12"},
			functionNames: []string{"A", "B", "C"},
			mainRoutine:   "A,B,A,C,B,C,B,A,C,B",
			functions: map[string]string{
				"A": "L,6,R,8,R,12,L,6,L,8",
				"B": "L,10,L,8,R,12",
				"C": "L,8,L,10,L,6,L,6",
			},
		},
	}
}
