package main

import (
	"testing"

	"github.com/mfesenko/adventofcode/2019/math"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func withTestWires(t *testing.T, test func([]wire)) {
	wires, err := parseWires([]string{"R8,U5,L5,D3", "U7,R6,D4,L4"})
	require.NoError(t, err)
	test(wires)
}

func TestFindBestIntersectionManhattan(t *testing.T) {
	withTestWires(t, func(wires []wire) {
		intersection := findBestIntersection(wires, compareManhattanDistance)

		assert.Equal(t, math.NewPoint(3, 3), intersection.point)
		assert.Equal(t, int64(6), intersection.manhattanDistance)
		assert.Equal(t, int64(40), intersection.stepCount)
	})
}

func TestFindBestIntersectionStepCount(t *testing.T) {
	withTestWires(t, func(wires []wire) {
		intersection := findBestIntersection(wires, compareStepCount)

		assert.Equal(t, math.NewPoint(6, 5), intersection.point)
		assert.Equal(t, int64(11), intersection.manhattanDistance)
		assert.Equal(t, int64(30), intersection.stepCount)
	})
}

func TestFindAllIntersections(t *testing.T) {
	withTestWires(t, func(wires []wire) {
		intersections := findAllIntersections(wires[0], wires[1])

		require.Equal(t, 2, len(intersections))
		assert.Contains(t, intersections, newIntersection(math.NewPoint(3, 3), 40))
		assert.Contains(t, intersections, newIntersection(math.NewPoint(6, 5), 30))
	})
}
