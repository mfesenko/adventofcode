package drawing

import (
	"fmt"
	"sync"

	"github.com/mfesenko/adventofcode/2019/async"
	"github.com/mfesenko/adventofcode/2019/math"
)

// Direction describes the direction in which robot is moving
type Direction int64

const (
	// Up represents up direction
	Up Direction = iota
	// Right represents right direction
	Right
	// Down represents down direction
	Down
	// Left represents left direction
	Left
)

// DirectionCommand describes the command for changing direction
type DirectionCommand int64

const (
	// TurnLeft represents turn left
	TurnLeft DirectionCommand = 0
	// TurnRight represents turn right
	TurnRight DirectionCommand = 1
)

const _unsupportedDirectionCommand = "unsupported direction command: %v"

// Computer represents a computer that provides instructions to the robot
type Computer interface {
	Execute()
	Input() chan int64
	Output() chan int64
}

// Robot represents a painting robot
type Robot struct {
	computer  Computer
	direction Direction
	position  math.Point
	covered   map[math.Point]*tile

	minX int64
	minY int64
	maxX int64
	maxY int64
}

type tile struct {
	color uint16
	count int
}

// NewRobot creates a robot
func NewRobot(computer Computer) *Robot {
	return &Robot{
		computer:  computer,
		direction: Up,
		position:  math.NewPoint(0, 0),
		covered:   map[math.Point]*tile{},
	}
}

// ChangeColor sets a color of a current tile
func (r *Robot) ChangeColor(color uint16) {
	curTile, ok := r.covered[r.position]
	if !ok {
		curTile = &tile{}
		r.covered[r.position] = curTile
	}
	curTile.color = color
	curTile.count++
}

// CoveredTiles returns the amount of tiles covered during a run of a robot
func (r *Robot) CoveredTiles() int {
	return len(r.covered)
}

// Run starts a robot run
func (r *Robot) Run() {
	done := &sync.WaitGroup{}
	done.Add(1)
	executor := async.NewExecutor(r.computer)
	executor.ExecuteAsync(done)
	for executor.Running() {
		r.computer.Input() <- int64(r.currentColor())
		color := uint16(<-r.computer.Output())
		command := DirectionCommand(<-r.computer.Output())
		r.ChangeColor(color)
		r.changeDirection(command)
		r.move()
	}
	done.Wait()
}

func (r *Robot) currentColor() uint16 {
	if curTile, ok := r.covered[r.position]; ok {
		return curTile.color
	}
	return Black
}

func (r *Robot) changeDirection(command DirectionCommand) {
	switch command {
	case TurnRight:
		r.direction = (r.direction + 1) % 4
	case TurnLeft:
		r.direction = (r.direction + 3) % 4
	default:
		panic(fmt.Sprintf(_unsupportedDirectionCommand, command))
	}
}

func (r *Robot) move() {
	switch r.direction {
	case Up:
		r.position = math.NewPoint(r.position.X, r.position.Y+1)
	case Down:
		r.position = math.NewPoint(r.position.X, r.position.Y-1)
	case Right:
		r.position = math.NewPoint(r.position.X+1, r.position.Y)
	case Left:
		r.position = math.NewPoint(r.position.X-1, r.position.Y)
	}
	r.updateDimensions()
}

func (r *Robot) updateDimensions() {
	r.minX = math.Min(r.minX, r.position.X)
	r.minY = math.Min(r.minY, r.position.Y)
	r.maxX = math.Max(r.maxX, r.position.X)
	r.maxY = math.Max(r.maxY, r.position.Y)
}

// Image converts the result of the program into an image
func (r *Robot) Image() *Image {
	result := NewImage(r.width(), r.height(), 1)
	for position, tile := range r.covered {
		if tile.color == White {
			y := position.Y - r.minY
			if r.minY < 0 {
				y = r.maxY - position.Y
			}

			x := position.X - r.minX
			if r.minX < 0 {
				x = r.maxX - position.X
			}

			result.SetPixel(int(x), int(y), 0, White)
		}
	}
	return result
}

func (r *Robot) height() int {
	return int(r.maxY - r.minY + 1)
}

func (r *Robot) width() int {
	return int(r.maxX - r.minX + 1)
}
