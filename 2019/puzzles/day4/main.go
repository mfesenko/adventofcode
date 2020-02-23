package main

import (
	"fmt"
)

func main() {
	start := int32(273025)
	end := int32(767253)
	fmt.Printf("Part 1: %v\n", findPasswordCount(start, end, twoAdjacentDigitsAreTheSame))
	fmt.Printf("Part 2: %v\n", findPasswordCount(start, end, exactlyTwoAdjacentDigitsAreTheSame))
}

func findPasswordCount(start int32, end int32, adjacencyRule rule) int {
	rule := andRule(
		isSixDigitNumber,
		isWithinRangeRule(start, end),
		isNeverDecreasing,
		adjacencyRule,
	)
	result := 0
	for i := start; i <= end; i++ {
		if rule(i) {
			result++
		}
	}
	return result
}
