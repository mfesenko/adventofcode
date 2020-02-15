package intcode

import (
	"math/rand"
	"testing"

	"github.com/Flaque/filet"
	"github.com/mfesenko/adventofcode/2019/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestWhenFailedToLoadDataFromFileThenLoadProgramReturnsAnError(t *testing.T) {
	testhelpers.WithTmpDir(t, func(dir string) {
		program, err := LoadProgram(dir)

		assert.Error(t, err)
		assert.Nil(t, program)
	})
}

func TestWhenFailedToParseProgramDataThenLoadProgramReturnsAnError(t *testing.T) {
	testhelpers.WithTmpDir(t, func(dir string) {
		file := filet.TmpFile(t, dir, "qwerty")

		program, err := LoadProgram(file.Name())

		assert.Error(t, err)
		assert.Nil(t, program)
	})
}

func TestWhenParsedProgramDataSuccessfullyThenLoadProgramReturnsAProgram(t *testing.T) {
	testhelpers.WithTmpDir(t, func(dir string) {
		program := randomProgram()
		file := filet.TmpFile(t, dir, program.AsString())

		loadedProgram, err := LoadProgram(file.Name())

		assert.NoError(t, err)
		assert.Equal(t, program, loadedProgram)
	})
}

func TestReadPanicsGivenNegativeAddress(t *testing.T) {
	program := randomProgram()

	assert.PanicsWithValue(t, _addressCanNotBeNegative, func() {
		program.Read(-1)
	})
}

func TestReadReturnsValueForAddressGivenValidAddress(t *testing.T) {
	program := randomProgram()
	address := int64(1)

	assert.Equal(t, program.code[address], program.Read(address))
}

func TestReadReturnsZeroGivenAddressHigherThenLen(t *testing.T) {
	program := randomProgram()
	address := program.Len() + 2

	assert.Equal(t, int64(0), program.Read(address))
}

func TestWritePanicsGivenNegativeAddress(t *testing.T) {
	program := randomProgram()

	assert.PanicsWithValue(t, _addressCanNotBeNegative, func() {
		program.Write(-1, randomValueBelowHundred())
	})
}

func TestWriteUpdatesValueForAddressGivenAddressLowerThenLen(t *testing.T) {
	program := randomProgram()
	startLen := program.Len()
	address := int64(1)
	value := randomValueBelowHundred()

	program.Write(address, value)
	assert.Equal(t, value, program.Read(address))
	assert.Equal(t, startLen, program.Len())
}

func TestWriteUpdatesValueForAddressGivenAddressHigherThenLen(t *testing.T) {
	program := randomProgram()
	value := randomValueBelowHundred()
	address := program.Len() + 2

	program.Write(address, value)
	assert.Equal(t, value, program.Read(address))
	assert.Equal(t, address+1, program.Len())
}

func TestLenReturnsLengthOfTheProgram(t *testing.T) {
	program := randomProgram()

	assert.Equal(t, int64(len(program.code)), program.Len())
}

func TestSetNounWritesValueToAddressOne(t *testing.T) {
	program := randomProgram()
	noun := randomValueBelowHundred()

	program.SetNoun(noun)

	assert.Equal(t, noun, program.code[1])
}

func TestSetVerbWritesValueToAddressTwo(t *testing.T) {
	program := randomProgram()
	verb := randomValueBelowHundred()

	program.SetVerb(verb)

	assert.Equal(t, verb, program.code[2])
}

func randomProgram() *Program {
	return NewProgram([]int64{
		randomValueAboveHundred(),
		randomValueAboveHundred(),
		randomValueAboveHundred(),
	})
}

func randomValueBelowHundred() int64 {
	return rand.Int63n(100)
}

func randomValueAboveHundred() int64 {
	return 101 + randomValueBelowHundred()
}
