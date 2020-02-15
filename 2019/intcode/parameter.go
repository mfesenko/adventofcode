package intcode

// ParameterMode defines how a parameter of an instruction is handled
type ParameterMode int

const (
	// PositionMode causes the parameter to be interpreted as a position. This is the default ParameterMode.
	PositionMode = ParameterMode(0)
	// ImmediateMode causes the parameter to be interpreted as a value.
	ImmediateMode = ParameterMode(1)
	// RelativeMode causes the parameter to be interpreted as a position plus the current relative base.
	RelativeMode = ParameterMode(2)
)
