package main

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"

	"github.com/Flaque/filet"
	"github.com/stretchr/testify/assert"
)

func TestWhenFailedToLoadDataFromFileThenReadModuleMassesReturnsAnError(t *testing.T) {
	withTmpDir(t, func(dir string) {
		masses, err := readModuleMasses(dir)

		assert.Error(t, err)
		assert.Nil(t, masses)
	})
}

func TestWhenFailedToConvertDataFromTheFileToIntThenReadModuleMassesReturnsAnError(t *testing.T) {
	withTmpDir(t, func(dir string) {
		file := filet.TmpFile(t, dir, "qwerty")

		masses, err := readModuleMasses(file.Name())

		assert.Error(t, err)
		assert.Nil(t, masses)
	})
}

func TestWhenLoadedInputDataThenReadModuleMassesReturnsMasses(t *testing.T) {
	withTmpDir(t, func(dir string) {
		testMasses := []int64{rand.Int63(), rand.Int63(), rand.Int63()}
		file := filet.TmpFile(t, dir, massesToString(testMasses))

		masses, err := readModuleMasses(file.Name())

		assert.NoError(t, err)
		assert.Equal(t, testMasses, masses)
	})
}

func massesToString(masses []int64) string {
	lines := make([]string, len(masses))
	for i, mass := range masses {
		lines[i] = strconv.FormatInt(mass, 10)
	}
	return strings.Join(lines, "\n")
}

func withTmpDir(t *testing.T, test func(string)) {
	dir := filet.TmpDir(t, "")
	defer filet.CleanUp(t)
	test(dir)
}

func TestCalculateFuelMass(t *testing.T) {
	masses := []int64{rand.Int63(), rand.Int63()}
	expectedMass := int64(0)
	for _, mass := range masses {
		expectedMass += mass * 2
	}

	fuelMass := calculateFuelMass(masses, func(mass int64) int64 {
		return mass * 2
	})

	assert.Equal(t, expectedMass, fuelMass)
}

func TestNaiveFuelCalculator(t *testing.T) {
	tests := map[int64]int64{
		12:     2,
		14:     2,
		1969:   654,
		100756: 33583,
	}
	for moduleMass, fuelMass := range tests {
		assert.Equal(t, fuelMass, naiveFuelCalculator(moduleMass))
	}
}

func TestSmartFuelCalculator(t *testing.T) {
	tests := map[int64]int64{
		12:     2,
		14:     2,
		1969:   966,
		100756: 50346,
	}
	for moduleMass, fuelMass := range tests {
		assert.Equal(t, fuelMass, smartFuelCalculator(moduleMass))
	}
}
