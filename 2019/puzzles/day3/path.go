package main

import (
	"errors"
	"strconv"
)

type direction rune

const (
	_up    = direction('U')
	_down  = direction('D')
	_left  = direction('L')
	_right = direction('R')
)

type path struct {
	dx int32
	dy int32
}

func parsePath(input string) (path, error) {
	value, err := strconv.ParseInt(input[1:], 10, 32)
	if err != nil {
		return path{}, err
	}

	diff := int32(value)

	var dx, dy int32
	direction := direction(input[0])
	switch direction {
	case _up:
		dy = diff
	case _down:
		dy = -diff
	case _right:
		dx = diff
	case _left:
		dx = -diff
	default:
		return path{}, errors.New("invalid direction")
	}

	return path{
		dx: dx,
		dy: dy,
	}, nil
}
