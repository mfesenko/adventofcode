package arcade

// JoystickPosition represents a position of the joystick
type JoystickPosition int

const (
	// JoystickNeutral means that the joystick is in the neutral position
	JoystickNeutral JoystickPosition = 0
	// JoystickLeft means that the joystick is tilted to the left
	JoystickLeft JoystickPosition = -1
	// JoystickRight means that the joystick is tilted to the right
	JoystickRight JoystickPosition = 1
)
