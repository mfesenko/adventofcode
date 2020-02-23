package main

import (
	"strings"

	"github.com/mfesenko/adventofcode/2019/math"
)

type wire []math.Interval

func parseWires(input []string) ([]wire, error) {
	wires := make([]wire, len(input))
	for i, line := range input {
		wire, err := parseWire(line)
		if err != nil {
			return nil, err
		}
		wires[i] = wire
	}
	return wires, nil
}

func parseWire(input string) (wire, error) {
	prevPoint := math.NewPoint(0, 0)
	w := wire{}
	for _, pathStr := range strings.Split(input, ",") {
		path, err := parsePath(pathStr)
		if err != nil {
			return nil, err
		}

		nextPoint := math.NewPoint(prevPoint.X+path.dx, prevPoint.Y+path.dy)
		w = append(w, math.NewInterval(prevPoint, nextPoint))
		prevPoint = nextPoint
	}
	return w, nil
}
