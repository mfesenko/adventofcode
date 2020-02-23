package main

import (
	"sort"

	"github.com/mfesenko/adventofcode/2019/math"
)

type (
	intersectionComparator func(intersection, intersection) int32

	intersection struct {
		point             math.Point
		manhattanDistance int32
		stepCount         int32
		cost              int32
	}
)

var _centralPort = math.NewPoint(0, 0)

func newIntersection(point math.Point, stepCount int32) intersection {
	return intersection{
		point:             point,
		manhattanDistance: point.ManhattanDistance(_centralPort),
		stepCount:         stepCount,
	}
}

func compareStepCount(a intersection, b intersection) int32 {
	return a.stepCount - b.stepCount
}

func compareManhattanDistance(a intersection, b intersection) int32 {
	return a.manhattanDistance - b.manhattanDistance
}

func findBestIntersection(wires []wire, comparator intersectionComparator) intersection {
	intersections := findAllIntersections(wires[0], wires[1])
	sort.SliceStable(intersections, func(i, j int) bool {
		return comparator(intersections[i], intersections[j]) < 0
	})
	return intersections[0]
}

func findAllIntersections(a wire, b wire) []intersection {
	intersectionMap := map[math.Point]int32{}
	stepsA := int32(0)
	for _, intervalA := range a {
		stepsB := int32(0)
		for _, intervalB := range b {
			point, ok := intervalA.FindIntersection(intervalB)
			if ok && _centralPort != point {
				totalSteps := stepsA + stepsB + intervalA.DistanceTo(point) + intervalB.DistanceTo(point)
				if prevSteps, ok := intersectionMap[point]; ok {
					intersectionMap[point] = math.Min(prevSteps, totalSteps)
				} else {
					intersectionMap[point] = totalSteps
				}
			}
			stepsB += intervalB.Length()
		}
		stepsA += intervalA.Length()
	}

	intersections := make([]intersection, 0, len(intersectionMap))
	for point, stepCount := range intersectionMap {
		intersections = append(intersections, newIntersection(point, stepCount))
	}
	return intersections
}
