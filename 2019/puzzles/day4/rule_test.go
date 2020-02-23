package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAndRule(t *testing.T) {
	trueRule := func(int32) bool {
		return true
	}

	falseRule := func(int32) bool {
		return false
	}

	assert.True(t, andRule(trueRule, trueRule)(222))
	assert.False(t, andRule(falseRule, trueRule)(222))
	assert.False(t, andRule(falseRule, falseRule)(222))
}

func TestIsSixDigitNumber(t *testing.T) {
	assert.True(t, isSixDigitNumber(223456))
	assert.False(t, isSixDigitNumber(23456))
}

func TestIsWithinRangeRule(t *testing.T) {
	rule := isWithinRangeRule(111, 999)
	assert.True(t, rule(222))
	assert.False(t, rule(100))
}

func TestTwoAdjacentDigitsAreTheSame(t *testing.T) {
	assert.True(t, twoAdjacentDigitsAreTheSame(223456))
	assert.False(t, twoAdjacentDigitsAreTheSame(123456))
}

func TestIsNeverDecreasing(t *testing.T) {
	assert.True(t, isNeverDecreasing(223456))
	assert.False(t, isNeverDecreasing(223450))
}

func TestExactlyTwoAdjacentDigitsAreTheSame(t *testing.T) {
	assert.True(t, exactlyTwoAdjacentDigitsAreTheSame(112233))
	assert.True(t, exactlyTwoAdjacentDigitsAreTheSame(111122))
	assert.False(t, exactlyTwoAdjacentDigitsAreTheSame(123444))
	assert.False(t, exactlyTwoAdjacentDigitsAreTheSame(123456))
}
