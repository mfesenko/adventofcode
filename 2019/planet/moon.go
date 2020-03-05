package planet

import (
	"fmt"
)

// Moon represents a Jupiter moon
type Moon struct {
	position *MultiDimValue
	velocity *MultiDimValue
}

// NewMoon creates a new moon
func NewMoon(position *MultiDimValue, velocity *MultiDimValue) *Moon {
	return &Moon{
		position: position,
		velocity: velocity,
	}
}

// ApplyGravity updates velocity of the moons by applying gravity
func (m *Moon) ApplyGravity(otherMoon *Moon) {
	for _, key := range m.position.Dimensions() {
		value := m.position.Get(key)
		otherValue := otherMoon.position.Get(key)
		if value == otherValue {
			continue
		}

		delta := int64(1)
		otherDelta := int64(-1)
		if value > otherValue {
			delta, otherDelta = otherDelta, delta
		}

		m.velocity.Set(key, m.velocity.Get(key)+delta)
		otherMoon.velocity.Set(key, otherMoon.velocity.Get(key)+otherDelta)
	}
}

// ApplyVelocity updates position of the moon by applying velocity
func (m *Moon) ApplyVelocity() {
	for _, key := range m.velocity.Dimensions() {
		m.position.Set(key, m.position.Get(key)+m.velocity.Get(key))
	}
}

// Position returns position of the moon
func (m *Moon) Position() *MultiDimValue {
	return m.position
}

// Velocity returns velocity of the moon
func (m *Moon) Velocity() *MultiDimValue {
	return m.velocity
}

// TotalEnergy returns a total energy of the moon
func (m *Moon) TotalEnergy() int64 {
	return m.potentialEnergy() * m.kineticEnergy()
}

func (m *Moon) potentialEnergy() int64 {
	return m.position.AbsSum()
}

func (m *Moon) kineticEnergy() int64 {
	return m.velocity.AbsSum()
}

// String returns string representation of the moon
func (m *Moon) String() string {
	return fmt.Sprintf("pos=%v, vel=%v", m.position.String(), m.velocity.String())
}

// DeepCopy returns a deep copy of the moon
func (m *Moon) DeepCopy() *Moon {
	return NewMoon(
		m.Position().DeepCopy(),
		m.Velocity().DeepCopy(),
	)
}

// ParseMoons parses information about moons from string input
func ParseMoons(input []string) ([]*Moon, error) {
	moons := make([]*Moon, len(input))
	for i, line := range input {
		moon, err := parseMoon(line)
		if err != nil {
			return nil, err
		}
		moons[i] = moon
	}
	return moons, nil
}

func parseMoon(input string) (*Moon, error) {
	position, err := parseXYZValue(input)
	if err != nil {
		return nil, err
	}
	return NewMoon(position, newXYZValue(0, 0, 0)), nil
}

func copyMoons(moons []*Moon) []*Moon {
	copy := make([]*Moon, len(moons))
	for i, moon := range moons {
		copy[i] = moon.DeepCopy()
	}
	return copy
}
