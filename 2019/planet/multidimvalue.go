package planet

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mfesenko/adventofcode/2019/math"
)

// MultiDimValue represents a multi dimensional value
type MultiDimValue struct {
	data map[string]int64
}

// NewMultiDimValue creates a new value
func NewMultiDimValue() *MultiDimValue {
	return &MultiDimValue{
		data: map[string]int64{},
	}
}

// Get returns a value for a given dimension
func (v *MultiDimValue) Get(dimension string) int64 {
	return v.data[dimension]
}

// Set writes a value for a given dimension
func (v *MultiDimValue) Set(dimension string, value int64) {
	v.data[dimension] = value
}

// Dimensions returns all dimensions available for the value
func (v *MultiDimValue) Dimensions() []string {
	dimensions := make([]string, 0, len(v.data))
	for dimension := range v.data {
		dimensions = append(dimensions, dimension)
	}
	sort.Strings(dimensions)
	return dimensions
}

// AbsSum returns a sum of absolute value for all dimensions
func (v *MultiDimValue) AbsSum() int64 {
	sum := int64(0)
	for _, value := range v.data {
		sum += math.Abs(value)
	}
	return sum
}

// String returns a string representation of the value
func (v *MultiDimValue) String() string {
	dimensions := v.Dimensions()
	builder := &strings.Builder{}
	builder.WriteString("<")
	for i, dimension := range dimensions {
		if i != 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(fmt.Sprintf("%s=% d", dimension, v.data[dimension]))
	}
	builder.WriteString(">")
	return builder.String()
}

// DeepCopy returns a deep copy of the value
func (v *MultiDimValue) DeepCopy() *MultiDimValue {
	copy := NewMultiDimValue()
	for _, dimension := range v.Dimensions() {
		copy.Set(dimension, v.Get(dimension))
	}
	return copy
}

func parseXYZValue(input string) (*MultiDimValue, error) {
	var x, y, z int64
	_, err := fmt.Sscanf(input, "<x=%d, y=%d, z=%d>", &x, &y, &z)
	if err != nil {
		return nil, err
	}
	return newXYZValue(x, y, z), nil
}

func newXYZValue(x int64, y int64, z int64) *MultiDimValue {
	value := NewMultiDimValue()
	value.Set("x", x)
	value.Set("y", y)
	value.Set("z", z)
	return value
}
