package intcode

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParameterCount(t *testing.T) {
	parameterCounts := map[OperationType]int64{
		AddOperation:                3,
		MultiplyOperation:           3,
		ReadInputOperation:          1,
		WriteOutputOperation:        1,
		JumpIfTrueOperation:         2,
		JumpIfFalseOperation:        2,
		LessThanOperation:           3,
		EqualsOperation:             3,
		AdjustRelativeBaseOperation: 1,
		HaltOperation:               0,
	}
	for operationType, count := range parameterCounts {
		assert.Equal(t, count, operationType.ParameterCount())
	}
}

func TestParseOperation(t *testing.T) {
	code := int64(1202)
	address := rand.Int63()
	parameterModes := []ParameterMode{
		RelativeMode,
		ImmediateMode,
		PositionMode,
	}

	operation := ParseOperation(address, code)

	assert.Equal(t, MultiplyOperation, operation.operationType)
	assert.Equal(t, parameterModes, operation.parameterModes)
}
