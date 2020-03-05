package planet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseXYZValueReturnsAnErrorWhenStringHasInvalidFormat(t *testing.T) {
	value, err := parseXYZValue("asdf")

	assert.Error(t, err)
	assert.Nil(t, value)
}

func TestParseXYZValueReturnsAValueWhenStringHasValidFormat(t *testing.T) {
	value, err := parseXYZValue("<x=22, y=13, z=-27>")

	assert.NoError(t, err)
	assert.Equal(t, newXYZValue(22, 13, -27), value)
}

func TestGetReturnsValueForDimension(t *testing.T) {
	x := int64(-123)
	y := int64(456)
	z := int64(789)
	value := newXYZValue(x, y, z)

	assert.Equal(t, x, value.Get("x"))
	assert.Equal(t, y, value.Get("y"))
	assert.Equal(t, z, value.Get("z"))
	assert.Equal(t, int64(0), value.Get("r"))
}

func TestSetWritesValueForDimension(t *testing.T) {
	value := NewMultiDimValue()
	dimension := "x"
	updatedValue := int64(45)
	assert.Equal(t, int64(0), value.Get(dimension))

	value.Set(dimension, updatedValue)

	assert.Equal(t, updatedValue, value.Get(dimension))
}

func TestDimensionsReturnsDimensionNames(t *testing.T) {
	value := NewMultiDimValue()
	assert.Empty(t, value.Dimensions())

	value.Set("x", 2)
	value.Set("y", 1)

	assert.Equal(t, []string{"x", "y"}, value.Dimensions())
}

func TestAbsSumReturnsSumOfAbsoluteValueForAllDimensions(t *testing.T) {
	value := newXYZValue(0, -2, 5)

	assert.Equal(t, int64(7), value.AbsSum())
}

func TestStringReturnsStringRepresentationOfTheValue(t *testing.T) {
	value := newXYZValue(-2, 56, 34)

	assert.Equal(t, "<x=-2, y= 56, z= 34>", value.String())
}

func TestDeepCopyReturnsACopyOfTheValue(t *testing.T) {
	value := newXYZValue(-2, 56, 34)

	copy := value.DeepCopy()

	assert.Equal(t, value, copy)
	value.Set("x", 22)
	assert.NotEqual(t, value, copy)
}
