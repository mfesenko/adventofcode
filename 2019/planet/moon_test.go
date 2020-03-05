package planet

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWhenGivenInvalidInputThenParseMoonsReturnsAnError(t *testing.T) {
	moons, err := ParseMoons([]string{"asdf"})

	assert.Error(t, err)
	assert.Nil(t, moons)
}

func TestWhenGivenValidInputThenParseMoonsReturnsMoons(t *testing.T) {
	moons, err := ParseMoons([]string{
		"<x=-1, y=3, z=2>",
		"<x=4, y=-8, z=2>",
	})

	assert.NoError(t, err)
	require.Equal(t, 2, len(moons))
	assert.Equal(t, newXYZValue(-1, 3, 2), moons[0].Position())
	assert.Equal(t, newXYZValue(0, 0, 0), moons[0].Velocity())
	assert.Equal(t, newXYZValue(4, -8, 2), moons[1].Position())
	assert.Equal(t, newXYZValue(0, 0, 0), moons[1].Velocity())
}

func TestApplyGravity(t *testing.T) {
	moon := NewMoon(newXYZValue(-1, 0, 2), newXYZValue(0, 0, 0))
	otherMoon := NewMoon(newXYZValue(4, -8, 2), newXYZValue(0, 0, 0))

	moon.ApplyGravity(otherMoon)

	assert.Equal(t, newXYZValue(-1, 0, 2), moon.Position())
	assert.Equal(t, newXYZValue(1, -1, 0), moon.Velocity())
	assert.Equal(t, newXYZValue(4, -8, 2), otherMoon.Position())
	assert.Equal(t, newXYZValue(-1, 1, 0), otherMoon.Velocity())
}

func TestApplyVelocity(t *testing.T) {
	moon := NewMoon(newXYZValue(-1, 3, 2), newXYZValue(4, -8, 0))

	moon.ApplyVelocity()

	assert.Equal(t, newXYZValue(3, -5, 2), moon.Position())
	assert.Equal(t, newXYZValue(4, -8, 0), moon.Velocity())
}

func TestTotalEnergyReturnsTotalEnergyOfTheMoon(t *testing.T) {
	moon := NewMoon(newXYZValue(-9, 3, 1), newXYZValue(-2, -4, -3))

	assert.Equal(t, int64(117), moon.TotalEnergy())
}

func TestStringReturnsStringRepresentationOfTheMoonState(t *testing.T) {
	moon := NewMoon(newXYZValue(-9, 3, 1), newXYZValue(-2, -4, -3))

	assert.Equal(t, "pos=<x=-9, y= 3, z= 1>, vel=<x=-2, y=-4, z=-3>", moon.String())
}

func TestDeepCopyReturnsACopyOfTheMoon(t *testing.T) {
	moon := NewMoon(newXYZValue(-9, 3, 1), newXYZValue(-2, -4, -3))

	copy := moon.DeepCopy()

	assert.Equal(t, moon, copy)
	moon.ApplyVelocity()
	assert.NotEqual(t, moon, copy)
}
