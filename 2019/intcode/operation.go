package intcode

// OperationType defines an operation supported by Intcode computer
type OperationType int

const (
	// AddOperation adds together numbers read from the first two parameters stores the result in a position defined by the third parameter
	AddOperation = OperationType(1)
	// MultiplyOperation multiplies numbers read from the first two parameters stores the result in a position defined by the third parameter
	MultiplyOperation = OperationType(2)
	// ReadInputOperation takes a single integer as input and saves it to the position given by its only parameter
	ReadInputOperation = OperationType(3)
	// WriteOutputOperation outputs the value of its only parameter
	WriteOutputOperation = OperationType(4)
	// JumpIfTrueOperation sets the instruction pointer to the value from the second parameter if the first parameter is non-zero. Otherwise, it does nothing.
	JumpIfTrueOperation = OperationType(5)
	// JumpIfFalseOperation sets the instruction pointer to the value from the second parameter if the first parameter is zero. Otherwise, it does nothing.
	JumpIfFalseOperation = OperationType(6)
	// LessThanOperation stores 1 in the position given by the third parameter if the first parameter is less than the second parameter. Otherwise, it stores 0.
	LessThanOperation = OperationType(7)
	// EqualsOperation stores 1 in the position given by the third parameter if the first parameter is equal to the second parameter. Otherwise, it stores 0.
	EqualsOperation = OperationType(8)
	// AdjustRelativeBaseOperation adjusts the relative base by the value of its only parameter.
	AdjustRelativeBaseOperation = OperationType(9)
	// HaltOperation stops execution of the program
	HaltOperation = OperationType(99)
)

// ParameterCount returns count of parameters for an OperationType
func (ot OperationType) ParameterCount() int64 {
	switch ot {
	case AddOperation:
		fallthrough
	case MultiplyOperation:
		fallthrough
	case LessThanOperation:
		fallthrough
	case EqualsOperation:
		return 3

	case ReadInputOperation:
		fallthrough
	case WriteOutputOperation:
		fallthrough
	case AdjustRelativeBaseOperation:
		return 1

	case JumpIfTrueOperation:
		fallthrough
	case JumpIfFalseOperation:
		return 2
	}

	return 0
}

// Operation represents an Intcode operation
type Operation struct {
	address        int64
	operationType  OperationType
	parameterModes []ParameterMode
}

// ParseOperation parses operation from the operation code
func ParseOperation(address int64, code int64) Operation {
	operationType := OperationType(code % 100)
	parameterCount := operationType.ParameterCount()
	parameterModes := make([]ParameterMode, parameterCount)
	div := code / 100
	for i := int64(0); i < parameterCount; i++ {
		parameterModes[i] = ParameterMode(div % 10)
		div /= 10
	}
	return Operation{
		address:        address,
		operationType:  operationType,
		parameterModes: parameterModes,
	}
}
