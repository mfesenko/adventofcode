package main

import (
	"errors"
	"strconv"
)

type direction rune

const (
	up    = direction('U')
	down  = direction('D')
	left  = direction('L')
	right = direction('R')
)

type path struct {
	dx int64
	dy int64
}

func parsePath(input string) (path, error) {
	diff, err := strconv.ParseInt(input[1:], 10, 32)
	if err != nil {
		return path{}, err
	}

	var dx, dy int64
	direction := direction(input[0])
	switch direction {
	case up:
		dy = diff
	case down:
		dy = -diff
	case right:
		dx = diff
	case left:
		dx = -diff
	default:
		return path{}, errors.New("invalid direction")
	}

	return path{
		dx: dx,
		dy: dy,
	}, nil
}
