package main

import (
	"fmt"
	"strings"
)

type rule func(int32) bool

func andRule(rules ...rule) rule {
	return func(password int32) bool {
		for _, r := range rules {
			if !r(password) {
				return false
			}
		}
		return true
	}
}

func isSixDigitNumber(password int32) bool {
	return password >= 100000 && password <= 999999
}

func isWithinRangeRule(start int32, end int32) rule {
	return func(password int32) bool {
		return password >= start && password <= end
	}
}

func twoAdjacentDigitsAreTheSame(password int32) bool {
	str := fmt.Sprint(password)
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			return true
		}
	}
	return false
}

func isNeverDecreasing(password int32) bool {
	str := fmt.Sprint(password)
	for i := 1; i < len(str); i++ {
		if str[i] < str[i-1] {
			return false
		}
	}
	return true
}

func exactlyTwoAdjacentDigitsAreTheSame(password int32) bool {
	str := fmt.Sprint(password)
	for i := 0; i <= 9; i++ {
		digit := fmt.Sprint(i)
		requiredWord := strings.Repeat(digit, 2)
		forbiddenWord := strings.Repeat(digit, 3)
		index := strings.Index(str, requiredWord)
		if index != -1 && strings.Index(str, requiredWord) != strings.Index(str, forbiddenWord) {

			return true
		}
	}
	return false
}
